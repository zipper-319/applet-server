package s3

import (
	"applet-server/internal/conf"
	"applet-server/internal/pkg/log"
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"io"
	"time"
)

type S3Service struct {
	*s3.S3
	*log.MyLogger
	Timeout int64
	Bucket  string
}

func NewS3Service(config *conf.Data, logger *log.MyLogger) (*S3Service, error) {
	logger.Infof("endpoint:%s;region:%s, access key:%s, secret key:%s", config.S3.Endpoint, config.S3.Region, config.S3.AccessKey, config.S3.SecretKey)
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	service := s3.New(sess, &aws.Config{
		DisableSSL: aws.Bool(false),
		Endpoint:   aws.String(config.S3.Endpoint),
		Region:     aws.String(config.S3.Region),
		Credentials: credentials.NewStaticCredentials(config.S3.AccessKey, config.S3.SecretKey, ""),
	})

	return &S3Service{
		S3:       service,
		MyLogger: logger,
		Timeout:  int64(config.S3.Timeout),
		Bucket:   config.S3.Bucket,
	}, nil
}

func (s *S3Service) Uploading(ctx context.Context, data []byte, filePath string) (err error) {

	ctxWithTimeOut, cancel := context.WithTimeout(ctx, time.Duration(s.Timeout)*time.Second)
	// Ensure the context is canceled to prevent leaking.
	// See context package for more information, https://golang.org/pkg/context/
	defer cancel()

	s.Debug("start time: ", time.Now())
	_, err = s.PutObjectWithContext(ctxWithTimeOut, &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filePath),
		Body:   bytes.NewReader(data),
	})
	s.Debug("end time ", time.Now())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			// If the SDK can determine the request or retry delay was canceled
			// by a context the CanceledErrorCode error code will be returned.
			err = errors.Wrap(err, "upload canceled due to timeout")
		} else {
			err = errors.Wrap(err, "failed to upload object")
		}
	}
	return err
}

func (s *S3Service) Download(ctx context.Context, filePath string) ([]byte, error) {
	ctxWithTimeOut, cancel := context.WithTimeout(ctx, time.Duration(s.Timeout)*time.Second)
	// Ensure the context is canceled to prevent leaking.
	// See context package for more information, https://golang.org/pkg/context/
	defer cancel()

	out, err := s.GetObjectWithContext(ctxWithTimeOut, &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filePath),
	})
	if err != nil {
		return nil, err
	}

	defer out.Body.Close()
	var result []byte
	for {
		temp := make([]byte, 1024)
		n, err := out.Body.Read(temp)
		if n > 0 {
			result = append(result, temp[:n]...)
		}
		if err == io.EOF || n == 0 {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (s *S3Service) GetFileList(bucketName, filePath string, timeout int) ([]string, error) {
	var objKeys []string

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	out, err := s.ListObjectsWithContext(ctx, &s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(filePath),
	})
	if err != nil {
		return nil, err
	}

	for _, content := range out.Contents {
		objKeys = append(objKeys, aws.StringValue(content.Key))
	}
	return objKeys, nil
}

func (s *S3Service) GetFilesAndCallback(bucketName, filePath string, timeout int, fun func(v *string) error) error {
	var objKeys []string

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	err := s.ListObjectsPagesWithContext(ctx, &s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(filePath),
	}, func(output *s3.ListObjectsOutput, b bool) bool {
		for _, content := range output.Contents {
			objKeys = append(objKeys, aws.StringValue(content.Key))
			fun(content.Key)
		}
		return true
	})
	return err
}
