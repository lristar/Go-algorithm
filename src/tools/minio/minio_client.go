package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func main() {
	GetObject()
}

func GetObject(){
	s3Client, err := minio.New("121.43.156.172:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("8Q1JY0B1JU2OIU9V33U3", "DcMU9+nlxxooSiVJ1v60u6fVfNI0e5gOcAPbZIaZ", ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if err := s3Client.FGetObject(context.Background(), "lzy-test", "aaabbb", "/home/lzy/test/aaabbb", minio.GetObjectOptions{}); err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully saved my-filename.csv")
}

func PutObject(){
	ctx := context.Background()
	endpoint := "121.43.156.172:9000"
	accessKeyID := "8Q1JY0B1JU2OIU9V33U3"
	secretAccessKey := "DcMU9+nlxxooSiVJ1v60u6fVfNI0e5gOcAPbZIaZ"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	bucketName := "lzy-test"
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the zip file
	objectName := "aaabbb"
	filePath := "/home/lzy/test/aaabbb"
	contentType := "application/zip"

	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}





