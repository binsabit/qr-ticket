package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"log"
)

var s3Storage *S3Manager

type S3Manager struct {
	downloader *s3manager.Downloader
	uploader   *s3manager.Uploader
	key        string
	bucket     string
}

func NewUploader(key, bucket string) {
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)
	downloader := s3manager.NewDownloader(sess)

	s3Storage = &S3Manager{
		downloader: downloader,
		uploader:   uploader,
		key:        key,
		bucket:     bucket,
	}
}

func (s S3Manager) Upload(file io.Reader) error {

	data := &s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(s.key),
		Body:   file,
	}
	_, err := s.uploader.Upload(data)
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	log.Println("successfully uploaded file")
	return nil
}

func (s S3Manager) Download(file io.WriterAt) error {
	n, err := s.downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(s.key),
	})
	if err != nil {
		return fmt.Errorf("failed to download file, %v", err)
	}
	fmt.Printf("file downloaded, %d bytes\n", n)
	return nil
}
