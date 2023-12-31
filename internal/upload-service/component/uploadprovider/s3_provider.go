package uploadprovider

import (
	"bytes"
	"common"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Provider struct {
	bucketName string
	region     string
	accessKey  string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketName, region, accessKey, secret, domain string) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		accessKey:  accessKey,
		secret:     secret,
		domain:     domain,
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			accessKey,
			secret,
			"",
		),
	})

	if err != nil {
		log.Fatalln(err)
	}

	provider.session = s3Session

	return provider
}

func (provider *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileTypes := http.DetectContentType(data)

	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileTypes),
		Body:        fileBytes,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	img := &common.Image{
		Url:       fmt.Sprintf("%s%s", provider.domain, dst),
		CloudName: "s3",
	}
	return img, nil
}
