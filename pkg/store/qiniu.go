package store

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"

	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Qiniu struct {
	mac           *qbox.Mac
	cfg           *storage.Config
	endpoint      string
	bucketManager *storage.BucketManager
	bucketName    string
}

func NewQiniu(bucketName, endpoint, accessKey, secretKey string, useSSL bool) *Qiniu {
	q := new(Qiniu)
	q.mac = qbox.NewMac(accessKey, secretKey)
	q.cfg = &storage.Config{
		UseHTTPS: useSSL,
	}
	q.endpoint = endpoint
	q.bucketName = bucketName
	q.bucketManager = storage.NewBucketManager(q.mac, q.cfg)
	return q
}

type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

func (q *Qiniu) Upload(path string, fileBytes []byte) error {
	var putPolicy = storage.PutPolicy{
		Scope:      fmt.Sprintf("%s:%s", q.bucketName, path),
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	putPolicy.Expires = 7200
	upToken := putPolicy.UploadToken(q.mac)
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	return formUploader.Put(nil, &ret, upToken, path, bytes.NewReader(fileBytes), int64(len(fileBytes)), nil)
}

func (q *Qiniu) Get(path string) (fileBytes []byte, err error) {
	url := storage.MakePublicURL(q.endpoint, path)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	fileBytes, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return
}

func (q *Qiniu) Delete(path string) error {
	return q.bucketManager.Delete(q.bucketName, path)
}

func (q *Qiniu) List(prefix, marker string, limit int) (fileNames []string, err error)  {
	entries, _, _, _, err := q.bucketManager.ListFiles(q.bucketName, prefix, "", marker, limit)
	if err != nil {
		return
	}
	for _, entry := range entries {
		fileNames = append(fileNames, entry.Key)
	}
	return
}
