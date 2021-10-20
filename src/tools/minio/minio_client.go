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

	if err := s3Client.FGetObject(context.Background(), "lzy-test", "release.asc", "/home/lzy/test/test-minio", minio.GetObjectOptions{}); err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully saved my-filename.csv")
}



