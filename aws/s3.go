package aws

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	s3KeyConst    = "S3_ACCESS_KEY"
	s3SecretConst = "S3_ACCESS_SECRET"
)

var S3Client *s3.Client

func InitS3() {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	s3Key := os.Getenv(s3KeyConst)
	s3Secret := os.Getenv(s3SecretConst)
	creds := credentials.NewStaticCredentialsProvider(s3Key, s3Secret, "")
	cfg.Credentials = creds
	S3Client = s3.NewFromConfig(cfg)
}

func ListS3Files() {
	ctx := context.Background()
	result, err := S3Client.ListBuckets(ctx, nil)
	if err != nil {
		log.Print("fail to list buckets")
	}

	for _, bucket := range result.Buckets {
		fmt.Println(*bucket.Name)
	}
}
