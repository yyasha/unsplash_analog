package cminio

import (
	"context"
	"log"
	"mime/multipart"
	"unsplash_analog/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Bucket names
var (
	Product_photos string = "productimgs"
	Product_models string = "productmodels"
)

var location string = "eu_ru"

var ctx context.Context = context.Background()

// MinioConnection func for opening minio connection.
func MinioConnection(bucketName string) (*minio.Client, error) {
	ctx := context.Background()
	useSSL := false
	// Initialize minio client object.
	minioClient, errInit := minio.New(config.Conf.MINIO_ADDR, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Conf.MINIO_ROOT_USER, config.Conf.MINIO_ROOT_PASSWORD, ""),
		Secure: useSSL,
	})
	if errInit != nil {
		return nil, errInit
	}
	// Check exists
	exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
	if errBucketExists != nil {
		return nil, errBucketExists
	}
	if !exists {
		// Create bucket
		err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			return nil, err
		} else {
			log.Printf("Successfully created %s\n", bucketName)
		}
	}
	return minioClient, errInit
}

// Upload file to bucket
func UploadFile(bucketName string, objectName string, fileBuffer *multipart.File, contentType string, fileSize int64) error {
	// Create minio connection.
	minioClient, err := MinioConnection(bucketName)
	if err != nil {
		return err
	}
	// Upload the zip file with PutObject
	info, err := minioClient.PutObject(ctx, bucketName, objectName, *fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	return nil
}

// Get file from bucket
func GetFile(bucketName string, objectName string) (*minio.Object, error) {
	// Create minio connection.
	minioClient, err := MinioConnection(bucketName)
	if err != nil {
		return &minio.Object{}, err
	}
	// Get object from minio
	minio_obj, err := minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{Checksum: true})
	return minio_obj, err
}

// Delete file from bucket
func DelFile(bucketName string, objectName string) error {
	// Create minio connection.
	minioClient, err := MinioConnection(bucketName)
	if err != nil {
		return err
	}
	err = minioClient.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{ForceDelete: true})
	return err
}
