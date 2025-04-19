package main

import (
	"context"
	"fmt"
	"os"

	"github.com/eyepipe/eye/internal/web"
	"github.com/eyepipe/eye/pkg/proto"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/pressly/goose/v3"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "eye",
		Copyright: proto.Copyright,
		Usage:     proto.Usage,
		Commands: []*cli.Command{
			{
				Name: "server",
				Flags: []cli.Flag{
					FlagConfigRequired,
					FlagAddress,
				},
				Action: func(ctx context.Context, c *cli.Command) (err error) {
					w := new(web.Web)
					server := NewServer()
					err = server.config.BindYAMLFile(c.String(FlagConfigRequired.Name))
					if err != nil {
						return fmt.Errorf("failed to bind config: %w", err)
					}

					s3Shards, err := server.GetS3Shards(ctx, server.config)
					if err != nil {
						return fmt.Errorf("failed to get S3 shards: %w", err)
					}

					storeShards, err := server.GetStoreShards(ctx, server.config)
					if err != nil {
						return fmt.Errorf("failed to get store shards: %w", err)
					}

					jwt, err := server.GetJWT(ctx, server.config)
					if err != nil {
						return fmt.Errorf("failed to get jwt %w", err)
					}

					app := fiber.New(fiber.Config{
						StreamRequestBody:            true,
						DisablePreParseMultipartForm: true,
						BodyLimit:                    server.config.ServerBodyLimitMiB << 20, // 4 Mb
					})
					app.Use(logger.New())

					w.SetS3(s3Shards)
					w.SetStores(storeShards)
					w.SetJwtToken(jwt)
					w.SetApp(app)
					w.SetupRoutesV1(app)

					err = app.Listen(c.String(FlagAddress.Name))
					if err != nil {
						return fmt.Errorf("failed to listen: %w", err)
					}

					return nil
				},
			},
			{
				Name: "db",
				Flags: []cli.Flag{
					FlagConfigRequired,
					FlagMigrationsDir,
				},
				Commands: []*cli.Command{
					{
						Name: "up",
						Action: func(ctx context.Context, c *cli.Command) (err error) {
							server := NewServer()
							err = server.config.BindYAMLFile(c.String(FlagConfigRequired.Name))
							if err != nil {
								return fmt.Errorf("failed to bind config: %w", err)
							}

							databases, err := SetupGooseDatabases(ctx, server)
							if err != nil {
								return fmt.Errorf("failed to setup goose databases: %w", err)
							}

							for _, db := range databases {
								if err = goose.UpContext(ctx, db, c.String(FlagMigrationsDir.Name)); err != nil {
									return fmt.Errorf("failed to goose.Up %w", err)
								}
							}

							return nil
						},
					},
					{
						Name: "down",
						Action: func(ctx context.Context, c *cli.Command) (err error) {
							server := NewServer()
							err = server.config.BindYAMLFile(c.String(FlagConfigRequired.Name))
							if err != nil {
								return fmt.Errorf("failed to bind config: %w", err)
							}

							databases, err := SetupGooseDatabases(ctx, server)
							if err != nil {
								return fmt.Errorf("failed to setup goose databases: %w", err)
							}

							for _, db := range databases {
								if err = goose.DownContext(ctx, db, c.String(FlagMigrationsDir.Name)); err != nil {
									return fmt.Errorf("failed to goose.Up %w", err)
								}
							}

							return nil
						},
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
