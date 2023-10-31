package utils

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

func UploadImageFile(file *multipart.FileHeader, client *s3.Client) (string, error) {

	bucket := viper.GetString("storage.BUCKET_NAME")

	src, _ := file.Open()

	generateRandomKey := uuid.NewV4().String()
	input := &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &generateRandomKey,
		Body:        src,
		ContentType: aws.String("image/png"),
	}

	_, err := client.PutObject(context.TODO(), input)

	if err != nil {
		return "", err
	}

	publicURL := fmt.Sprintf("https://pub-29bbc3f8a3274307ba076b469d80de0e.r2.dev/%s", generateRandomKey)

	return publicURL, nil

}
