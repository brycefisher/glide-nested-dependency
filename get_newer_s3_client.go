package glide_nested_dependency

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetNewerS3Client(region string, forcePathStyle bool) *s3.S3 {
	sess := session.New(&aws.Config{
		Region:           &region,
		S3ForcePathStyle: &forcePathStyle,
	})
	return s3.New(sess)
}
