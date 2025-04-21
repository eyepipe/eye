package main

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/eyepipe/eye/internal/lib/crypto2"
	"github.com/eyepipe/eye/internal/pkg/buildinfo"
	"github.com/eyepipe/eye/internal/pkg/container"
	"github.com/eyepipe/eye/internal/pkg/services/encryption_manager"
	"github.com/eyepipe/eye/internal/pkg/services/encryption_service"
	"github.com/eyepipe/eye/internal/pkg/services/export_service"
	"github.com/eyepipe/eye/internal/pkg/services/input_resolver_service"
	"github.com/eyepipe/eye/internal/pkg/services/keygen_service"
	"github.com/eyepipe/eye/internal/pkg/services/signature_reader_service"
	"github.com/eyepipe/eye/pkg/proto"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "eye",
		Version:   buildinfo.FullVersionString(),
		Copyright: proto.Copyright,
		Usage:     proto.Usage,
		Commands: []*cli.Command{
			{
				Name: "keygen",
				Flags: []cli.Flag{
					FlagScheme,
					FlagTimeout7s,
				},
				Description: "generate identity file with private keys",
				Action: func(ctx context.Context, c *cli.Command) error {
					ctx, cancel := context.WithTimeout(ctx, c.Duration(FlagTimeout7s.Name))
					defer cancel()

					service := input_resolver_service.NewService()
					scheme, err := service.ResolveSchemeURLOrJSON(ctx, c.String(FlagScheme.Name))
					if err != nil {
						return fmt.Errorf("failed to resolve the scheme: %w", err)
					}

					// TODO: add scheme validation service
					keygen := keygen_service.NewService(scheme)
					err = keygen.GenerateKeyAgreementPair()
					if err != nil {
						return err
					}

					err = keygen.GenerateSignerPair()
					if err != nil {
						return err
					}

					export := export_service.NewExportService(keygen.GetContainer())
					return export.ExportPrivate(c.Writer)
				},
			},
			{
				Name:        "public",
				Description: "generate identity file with private keys",
				Action: func(ctx context.Context, c *cli.Command) (err error) {
					ct, err := ImportContainer(ctx, c.Reader, "")
					if err != nil {
						return fmt.Errorf("failed to import container: %w", err)
					}

					export := export_service.NewExportService(ct)
					return export.ExportPublic(c.Writer)
				},
			},
			{
				Name: "encrypt",
				Flags: []cli.Flag{
					IFlag,
					PFlag,
				},
				Commands: []*cli.Command{
					{
						Name: "send",
						Flags: []cli.Flag{
							ContractURLFlag,
						},
						Action: func(ctx context.Context, c *cli.Command) (err error) {
							resolver := input_resolver_service.NewService()
							i, err := resolver.ResolveContainerReaderOrFileOrUrl(ctx, nil, c.String(IFlag.Name))
							p, err := resolver.ResolveContainerReaderOrFileOrUrl(ctx, nil, c.String(PFlag.Name))

							service := encryption_service.NewService(i, p)
							manager := encryption_manager.NewManager(service)

							res, err := manager.SendEncrypt(ctx, c.Reader, c.String(ContractURLFlag.Name))
							if err != nil {
								return fmt.Errorf("failed to manager.SendEncrypt: %w", err)
							}

							_, err = fmt.Fprintln(c.Writer, res.URL)
							return err
						},
					},
				},
				Action: func(ctx context.Context, c *cli.Command) (err error) {
					resolver := input_resolver_service.NewService()

					i, err := resolver.ResolveContainerReaderOrFileOrUrl(ctx, nil, c.String(IFlag.Name))
					if err != nil {
						return fmt.Errorf("failed to reolve identity: %w", err)
					}

					r, err := resolver.ResolveContainerReaderOrFileOrUrl(ctx, nil, c.String(PFlag.Name))
					if err != nil {
						return fmt.Errorf("failed to reolve participant: %w", err)
					}

					service := encryption_service.NewService(i, r)
					manager := encryption_manager.NewManager(service)

					sig, err := manager.Encrypt(ctx, c.Reader, c.Writer)
					if err != nil {
						return fmt.Errorf("failed to manager.Encrypt: %w", err)
					}

					_, err = c.ErrWriter.Write([]byte(fmt.Sprintf("%x", sig)))
					if err != nil {
						return fmt.Errorf("failed to write signature: %w", err)
					}

					return nil
				},
			},
			{
				Name: "decrypt",
				Flags: []cli.Flag{
					IFlag,
					PFlag,
					FlagSigHex,
					FlagSig,
				},
				Action: func(ctx context.Context, c *cli.Command) (err error) {
					i, err := ImportContainer(ctx, nil, c.String(IFlag.Name))
					p, err := ImportContainer(ctx, nil, c.String(PFlag.Name))

					service := encryption_service.NewService(i, p)
					manager := encryption_manager.NewManager(service)

					if url := c.Args().Get(0); len(url) > 0 {
						err = manager.DownloadDecrypt(ctx, c.Args().Get(0), c.Writer)
						switch {
						case err != nil:
							return fmt.Errorf("failed to download and decrypt: %w", err)
						default:
							return nil
						}
					}

					sig, err := signature_reader_service.NewService().ReadHex(ctx, signature_reader_service.ReadOpts{
						HEX:      c.String(FlagSigHex.Name),
						Filename: c.String(FlagSig.Name),
					})
					if err != nil {
						return fmt.Errorf("failed to read signature: %w", err)
					}
					err = manager.Decrypt(ctx, sig, c.Reader, c.Writer)
					if err != nil {
						return fmt.Errorf("failed to manager.Decrypt: %w", err)
					}

					return nil
				},
			},
			{
				Name: "sign",
				Flags: []cli.Flag{
					IFlag,
				},
				Action: func(ctx context.Context, c *cli.Command) (err error) {
					ct, err := ImportContainer(ctx, nil, c.String(IFlag.Name))
					if err != nil {
						return fmt.Errorf("failed to import container: %w", err)
					}

					signer := ct.Scheme.SignAlgo.ToSigner()
					hash := signer.GetHashier().New()
					_, err = io.Copy(hash, c.Reader)
					if err != nil {
						return fmt.Errorf("failed to copy: %w", err)
					}

					sig, err := signer.Sign(hash.Sum(nil), ct.SignerPair.Private)
					if err != nil {
						return fmt.Errorf("failed to sign: %w", err)
					}

					_, err = fmt.Fprint(c.Writer, hex.EncodeToString(sig))
					if err != nil {
						return fmt.Errorf("failed to write: %w", err)
					}

					return nil
				},
			},
			{
				Name: "verify",
				Flags: []cli.Flag{
					PFlag,
					FlagSigHex,
					FlagSig,
				},
				Action: func(ctx context.Context, c *cli.Command) (err error) {
					p, err := ImportContainer(ctx, nil, c.String(PFlag.Name))
					if err != nil {
						return fmt.Errorf("failed to import container: %w", err)
					}

					signatureReader := signature_reader_service.NewService()
					sig, err := signatureReader.ReadHex(ctx, signature_reader_service.ReadOpts{
						Filename: c.String(FlagSig.Name),
						HEX:      c.String(FlagSigHex.Name),
					})
					if err != nil {
						return fmt.Errorf("failed to read signature file: %w", err)
					}

					signer := p.Scheme.SignAlgo.ToSigner()
					hash := signer.GetHashier().New()
					_, err = io.Copy(hash, c.Reader)
					if err != nil {
						return fmt.Errorf("failed to copy: %w", err)
					}

					err = signer.Verify(hash.Sum(nil), sig, p.SignerPair.Public)
					switch {
					case errors.Is(err, crypto2.ErrSignatureInvalid):
						return fmt.Errorf("invalid signature ❌")
					case err != nil:
						return fmt.Errorf("failed to verify signature: %w", err)
					default:
						_, err = fmt.Fprint(c.Writer, "valid signature ✅\n")
						return err
					}
				},
			},
			{
				Name: "hex",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "dec",
						Aliases: []string{"decode"},
					},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					switch {
					case c.Bool("dec"):
						if _, err := io.Copy(c.Writer, hex.NewDecoder(c.Reader)); err != nil {
							return fmt.Errorf("failed to io.Copy: %w", err)
						}
					default:
						if _, err := io.Copy(hex.NewEncoder(c.Writer), c.Reader); err != nil {
							return fmt.Errorf("failed to io.Copy: %w", err)
						}
					}

					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ImportContainer(ctx context.Context, reader io.Reader, path string) (ct *container.Container, err error) {
	var bytea []byte
	if reader != nil {
		bytea, err = io.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("failed identity from io.ReadAll: %w", err)
		}
	} else {
		bytea, err = os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read identity file: %w", err)
		}
	}

	ct = container.NewEmptyContainer()
	service := export_service.NewExportService(ct)
	err = service.Import(bytea)
	if err != nil {
		return nil, err
	}

	return ct, nil
}
