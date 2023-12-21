package aws

import "github.com/aws/aws-sdk-go-v2/service/s3"

func ListS3Files() {
	opt := &s3.Options{}
	s3.New(*opt)
}
