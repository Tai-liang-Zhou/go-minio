package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	endpoint := "192.168.0.102:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	// ctx, cancel := context.WithCancel(context.Background())

	// defer cancel()

	// objectCh := minioClient.ListObjects(ctx, "kla", minio.ListObjectsOptions{
	// 	Prefix:    "",
	// 	Recursive: true,
	// })
	// for object := range objectCh {
	// 	if object.Err != nil {
	// 		fmt.Println(object.Err)
	// 		return
	// 	}
	// 	fmt.Println(object.Key)
	// }
	// log.Printf("%#v\n", minioClient) // minioClient is now set up

	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.png\"")

	presignedURL, err := minioClient.PresignedGetObject(context.Background(), "kla", "2022-04-02-01-57-16.png", time.Minute*60*24, reqParams)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Successfully generated presigned URL", presignedURL)
}
