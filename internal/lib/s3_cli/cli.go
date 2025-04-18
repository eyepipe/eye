package s3_cli

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/samber/lo"
	"github.com/shlima/oi/dsn"
)

type Cli struct {
	urn           string // unique identifier like "bucket@endpoint"
	bucket        string
	uploader      *manager.Uploader
	presignClient *s3.PresignClient
}

func New(url string) (*Cli, error) {
	parsed, err := dsn.Parse(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DSN: %w", err)
	}

	region, err := parsed.QueryE("region")
	if err != nil {
		return nil, fmt.Errorf("failed to extract region: %w", err)
	}

	urn, err := parsed.QueryE("urn")
	if err != nil {
		return nil, fmt.Errorf("failed to extract urn: %w", err)
	}

	bucket, err := parsed.QueryE("bucket")
	if err != nil {
		return nil, fmt.Errorf("failed to extract bucket: %w", err)
	}

	accessKey, err := parsed.QueryE("accessKey")
	if err != nil {
		return nil, fmt.Errorf("failed to extract accessKey: %w", err)
	}

	secretKey, err := parsed.QueryE("secretKey")
	if err != nil {
		return nil, fmt.Errorf("failed to extract secretKey: %w", err)
	}

	config := aws.Config{
		Region: region,
		//HTTPClient: &http.Client{
		//	Transport: &http.Transport{
		//		DialContext: (&net.Dialer{
		//			Timeout:   5 * time.Second, // Dial timeout
		//			KeepAlive: 30 * time.Second,
		//		}).DialContext,
		//		TLSHandshakeTimeout:   5 * time.Second,
		//		ResponseHeaderTimeout: 10 * time.Second,
		//		ExpectContinueTimeout: 1 * time.Second,
		//		IdleConnTimeout:       90 * time.Second,
		//		MaxIdleConns:          10,
		//	},
		//},
		BaseEndpoint: lo.ToPtr(parsed.Format(dsn.Scheme, dsn.Host)),
		Credentials: aws.NewCredentialsCache(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
	}

	client := s3.NewFromConfig(config)
	presignClient := s3.NewPresignClient(client)
	uploader := manager.NewUploader(client)

	return &Cli{
		urn:           urn,
		bucket:        bucket,
		uploader:      uploader,
		presignClient: presignClient,
	}, nil
}

// GetURN returns unique identifier
func (c *Cli) GetURN() string {
	return c.urn
}

// UploadReadAll blocks thread and read all from the reader
func (c *Cli) UploadReadAll(ctx context.Context, reader io.Reader, acl ACL, key string) (*UploadOutput, error) {
	return c.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
		ACL:    acl,
		Body:   reader,
	})
}

func (c *Cli) GetPresignedURL(ctx context.Context, key string, ttl time.Duration) (string, error) {
	req, err := c.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(ttl))
	switch {
	case err != nil:
		return "", err
	default:
		return req.URL, nil
	}
}
