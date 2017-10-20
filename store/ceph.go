package store

import (
	"github.com/minio/minio-go"

	"log"
	"bytes"
	"io/ioutil"
)

const (
	endpoint = "cs54:8080"
	accessKeyID = "Z8DFXNIGF71P8K71D0ZA"
	secretAccessKey = "VI0k5TecQOjwowWzhfapKRUfJfhs041hphieEhIM"
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
	_, err := c.minioClient.PutObject(bucketName, path, bytes.NewReader(fileBytes), int64(len(fileBytes)), minio.PutObjectOptions{})
	return err
}

func (c *Ceph) Get(path string) (fileBytes []byte, err error) {
	o, err := c.minioClient.GetObject(bucketName, path, minio.GetObjectOptions{})
	if err != nil {
		return
	}
	fileBytes, err = ioutil.ReadAll(o)
	o.Close()
	return
}

func (c *Ceph) Delete(path string) error {
	return c.minioClient.RemoveObject(bucketName, path)
}

func (c *Ceph) List(prefix string, limit int) (fileNames []string, err error) {
	doneCh := make(chan struct{})

	defer close(doneCh)

	objectCh := c.minioClient.ListObjectsV2(bucketName, prefix, true, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			break
		}
		fileNames = append(fileNames, object.Key)
	}
	return
}

