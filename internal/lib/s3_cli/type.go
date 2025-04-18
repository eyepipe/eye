package s3_cli

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type (
	ACL          = types.ObjectCannedACL
	UploadOutput = manager.UploadOutput
)

var (
	ACLPrivate                ACL = types.ObjectCannedACLPrivate
	ACLPublicRead             ACL = types.ObjectCannedACLPublicRead
	ACLPublicReadWrite        ACL = types.ObjectCannedACLPublicReadWrite
	ACLAuthenticatedRead      ACL = types.ObjectCannedACLAuthenticatedRead
	ACLAwsExecRead            ACL = types.ObjectCannedACLAwsExecRead
	ACLBucketOwnerRead        ACL = types.ObjectCannedACLBucketOwnerRead
	ACLBucketOwnerFullControl ACL = types.ObjectCannedACLBucketOwnerFullControl
)

type ICli interface {
	GetURN() string
	UploadReadAll(ctx context.Context, reader io.Reader, acl ACL, key string) (*UploadOutput, error)
	GetPresignedURL(ctx context.Context, key string, ttl time.Duration) (string, error)
}
