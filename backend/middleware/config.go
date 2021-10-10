package middleware

import (
	"os"
)

type Config struct {
	Aws struct {
		S3 struct {
			Region string
			Bucket string
			AccessKeyID string
			SecretAccessKey string
		}
	}
}

func NewConfig() *Config {
	c := new(Config)

	c.Aws.S3.Region = os.Getenv("AWS_REGION")
	c.Aws.S3.Bucket = os.Getenv("AWS_BUCKET")
	c.Aws.S3.AccessKeyID = os.Getenv("AWS_ACCESS_KEY")
	c.Aws.S3.SecretAccessKey = os.Getenv("AWS_SECRET_KEY")
	return c
}