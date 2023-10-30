package minio

import (
	"applet-server/internal/conf"
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type Client struct {
	*minio.Client
	bucket string
	region string
}

func NewMinioClient(config *conf.Data) *Client {
	client, err := minio.New(config.Minio.Addr, &minio.Options{
		Creds: credentials.NewStaticV4(config.Minio.AccessKey, config.Minio.SecretKey, ""),
	})
	if err != nil {
		panic(err)
	}
	return &Client{
		Client: client,
		bucket: config.Minio.Bucket,
		region: config.Minio.Region,
	}
}

func (c *Client) Uploading(ctx context.Context, video []byte, fileName string) (string, error) {
	file := bytes.NewReader(video)
	exists, err := c.BucketExists(ctx, c.bucket)
	if err != nil {
		return "", err
	}
	if !exists {

		err = c.MakeBucket(ctx, c.bucket, minio.MakeBucketOptions{Region: c.region})
		if err != nil {
			return "", err

		}
	}
	contentType := "video/pcm"

	info, err := c.PutObject(ctx, c.bucket, fileName, file, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", fileName, info.Size)

	return info.Location, nil

}
