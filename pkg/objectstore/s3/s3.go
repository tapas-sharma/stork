package s3

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/libopenstorage/secrets/aws/credentials"
	stork_api "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	"gocloud.dev/blob"
	"gocloud.dev/blob/s3blob"
)

func getSession(backupLocation *stork_api.BackupLocation) (*session.Session, error) {
	// AWS SDK fetches the correct endpoint based on region provided if endpoint is passed empty
	var endpoint string
	if backupLocation.Location.S3Config.Endpoint == "s3.amazonaws.com" {
		endpoint = ""
	} else {
		endpoint = backupLocation.Location.S3Config.Endpoint
	}
	awsCreds, err := credentials.NewAWSCredentials(
		backupLocation.Location.S3Config.AccessKeyID,
		backupLocation.Location.S3Config.SecretAccessKey,
		"",
		backupLocation.Location.S3Config.UseIam, // runningOnEc2 when set ec2 role credentials will be used
	)
	if err != nil {
		return nil, err
	}
	creds, err := awsCreds.Get()
	if err != nil {
		return nil, err
	}

	return session.NewSession(&aws.Config{
		Endpoint:         aws.String(endpoint),
		Credentials:      creds,
		Region:           aws.String(backupLocation.Location.S3Config.Region),
		DisableSSL:       aws.Bool(backupLocation.Location.S3Config.DisableSSL),
		S3ForcePathStyle: aws.Bool(true),
	})
}

// GetBucket gets a reference to the bucket for that backup location
func GetBucket(backupLocation *stork_api.BackupLocation) (*blob.Bucket, error) {
	sess, err := getSession(backupLocation)
	if err != nil {
		return nil, err
	}
	return s3blob.OpenBucket(context.Background(), sess, backupLocation.Location.Path, nil)
}

// CreateBucket creates a bucket for the bucket location
func CreateBucket(backupLocation *stork_api.BackupLocation) error {
	sess, err := getSession(backupLocation)
	if err != nil {
		return err
	}

	input := &s3.CreateBucketInput{
		Bucket: &backupLocation.Location.Path,
	}
	_, err = s3.New(sess).CreateBucket(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == s3.ErrCodeBucketAlreadyOwnedByYou {
				return nil
			}
		}
	}
	return err
}
