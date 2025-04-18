package main

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"

	"github.com/eyepipe/eye/internal/lib/jwt_token"
	"github.com/eyepipe/eye/internal/lib/pool"
	"github.com/eyepipe/eye/internal/lib/s3_cli"
	"github.com/eyepipe/eye/internal/lib/shard"
	"github.com/eyepipe/eye/internal/pkg/configuration"
	"github.com/eyepipe/eye/internal/pkg/store"
	"github.com/eyepipe/eye/internal/web"
)

type Server struct {
	config *configuration.Configuration
}

func NewServer() *Server {
	return &Server{
		config: configuration.New(),
	}
}

func (s *Server) GetS3Shards(ctx context.Context, config *configuration.Configuration) (*shard.Shards[s3_cli.ICli], error) {
	shards := shard.NewShards[s3_cli.ICli]()
	for _, dsn := range config.S3ShardDSN {
		s3, err := s3_cli.New(dsn)
		if err != nil {
			return nil, fmt.Errorf("failed to s3_cli.New: %w", err)
		}

		shards.Add(s3)
	}

	return shards, nil
}

func (s *Server) GetStoreShards(ctx context.Context, config *configuration.Configuration) (*shard.Shards[web.IStore], error) {
	shards := shard.NewShards[web.IStore]()
	for _, filename := range config.DDShardFiles {
		connections := make([]*sql.DB, 0, len(config.S3ShardDSN))
		for i := 0; i < config.DbPoolSize; i++ {
			db, err := store.Connect(filename)
			switch {
			case err != nil:
				return nil, fmt.Errorf("failed to store.Connect: %w", err)
			default:
				connections = append(connections, db)
			}
		}

		shards.Add(store.New(pool.New(connections)))
	}

	return shards, nil
}

func (s *Server) GetJWT(ctx context.Context, config *configuration.Configuration) (*jwt_token.Token, error) {
	algo, err := jwt_token.ParseAlgo(config.JwtAlgo)
	if err != nil {
		return nil, fmt.Errorf("failed to jwt.ParseAlgo: %w", err)
	}

	secret, err := hex.DecodeString(config.JwtSecretHex)
	switch {
	case err != nil:
		return nil, fmt.Errorf("failed to hex decode jwt secret: %w", err)
	case len(secret) == 0:
		return nil, fmt.Errorf("jwt secret has zero size")
	}

	token, err := jwt_token.New(config.JwtIss, algo, secret)
	if err != nil {
		return nil, fmt.Errorf("failed to jwt_token.New: %w", err)
	}

	return token, nil
}
