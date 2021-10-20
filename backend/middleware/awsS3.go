package middleware

import (
	"errors"
	"fmt"
	"mime/multipart"
	"bytes"
  "log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsS3 struct {
	Config *Config
	Keys AwsS3URLs
	Uploader *s3manager.Uploader
}

type AwsS3URLs struct {
	Path string
}

func NewAwsS3() *AwsS3 {
	config := NewConfig()
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
				AccessKeyID: config.Aws.S3.AccessKeyID,
        SecretAccessKey: config.Aws.S3.SecretAccessKey,
			}),
			Region: aws.String(config.Aws.S3.Region),
		},
	}))
	return &AwsS3{
		Config: config,
		Keys: AwsS3URLs{
			Path: "/images",
		},
		Uploader: s3manager.NewUploader(sess),
	}
}

func (a *AwsS3) Upload(file multipart.File, fileName string, extension string) (url string, err error) {
	if fileName == "" {
		return "", errors.New("fileName is required")
	}

	var contentType string
	switch extension {
	case "jpg":
		contentType = "image/jpeg"
	case "JPG":
		contentType = "image/jpeg"
	case "jpeg":
		contentType = "image/jpeg"
	case "JPEG":
		contentType = "image/jpeg"
	case "gif":
		contentType = "image/gif"
	case "GIF":
		contentType = "image/gif"
	case "png":
		contentType = "image/png"
	case "PNG":
		contentType = "image/png"
	default: return "", errors.New("this extension is invalid")
	}

	result, err := a.Uploader.Upload(&s3manager.UploadInput{
		ACL: aws.String("public-read"),
		Body: file,
		Bucket: aws.String(a.Config.Aws.S3.Bucket),
		// content-type の設定も忘れずに
		ContentType: aws.String(contentType),
		Key: aws.String(a.Keys.Path + "/" + fileName + "." + extension),
	})

	if err != nil {
			return "", fmt.Errorf("failed to upload file, %v", err)
	}

	return result.Location, nil
}


func DownloadConfigure(fileName string) {
	filePath := "../dist/images/" + fileName
	bucket := "nskobe"
	key := "images/" + fileName
	awsRegion := "ap-northeast-1"

	file, err := os.Create(filePath)
  if err != nil {
    log.Fatal(err)
  }
	newSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
  svc := s3.New(newSession, &aws.Config{
    Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
      AccessKeyID: os.Getenv("AWS_ACCESS_KEY"),
      SecretAccessKey: os.Getenv("AWS_SECRET_KEY"),
    }),
		Region: aws.String(awsRegion),
	})
	downloadKey := &s3.GetObjectInput{
		// Bucket ダウンロードするS3のバケット名
		Bucket: aws.String(bucket),
		// Key ダウンロードするオブジェクト名
		Key: aws.String(key),
	}
  image, err := svc.GetObject(downloadKey)
	if err != nil {
		log.Fatal(err)
	}
  buf := new(bytes.Buffer)
	buf.ReadFrom(image.Body)
  _, err = file.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("S3からダウンロードが完了しました。")
}

