package store

import (
	"github.com/minio/minio-go"
	"log"
	"strings"
)

const (
	endpoint = "cs54:8080"
	accessKeyID = "Z8DFXNIGF71P8K71D0ZA"
	secretAccessKey = "zuf+VI0k5TecQOjwowWzhfapKRUfJfhs041hphieEhIM"
	useSSL = false
	bucketName = "cephtest"
)

type Ceph struct {
	minioClient *minio.Client
}

func NewCeph() *Ceph {
	c := new(Ceph)
	var err error
	c.minioClient, err = minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)

	if err != nil {
		log.Fatal(err)
	}
	c.minioClient.MakeBucket(bucketName, "")
	if err != nil {
		exists, err := c.minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	return c
}

func (c *Ceph) Upload(path string, fileBytes []byte) error {
	pathes := strings.Split(path, "/")
	lastIndex := len(pathes) - 1
	fileName := pathes[lastIndex]
	_, err := c.minioClient.FPutObject(bucketName, fileName, path, minio.PutObjectOptions{})
	return err
}

func (c *Ceph) Get(path string) (fileBytes []byte, err error) {
	return nil, nil
}

func (c *Ceph) Delete(path string) error {
	return nil
}

func (c *Ceph) List(prefix string, limit int) (fileNames []string, err error) {
	return nil, nil
}

