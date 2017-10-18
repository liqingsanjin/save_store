package store

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
)

const (
	accessKey = "VjgieAVwG723rZomW6h0SRdEIXHe_vNnBDtDw527"
	secretKey = "bnQc-7eEPnnN5LEolwfxbAQjwr07d85CShh9lfLg"
	bucket    = "liqingsanjin"
	domain    = "http://oxx1uk7mm.bkt.clouddn.com"
)

type Qiniu struct {
	mac *qbox.Mac
	cfg *storage.Config
	bucketManager *storage.BucketManager
}

func NewQiniu() *Qiniu {
	q := new(Qiniu)
	q.mac = qbox.NewMac(accessKey, secretKey)
	q.cfg = &storage.Config{
		UseHTTPS: false,
	}
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
		Scope: fmt.Sprintf("%s:%s", bucket, path),
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
	url := storage.MakePublicURL(domain, path)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	fileBytes, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return
}

func (q *Qiniu) Delete(path string) error {
	return q.bucketManager.Delete(bucket, path)
}

func (q *Qiniu) List(prefix string, limit int) (fileNames []string, err error) {
	err = nil

	marker := ""
	delimiter := ""
	for {
		entries, _, nextMarker, hashNext, err := q.bucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)
		if err != nil {
			fmt.Println("list error,", err)
			break
		}
		for _, entry := range entries {
			fileNames = append(fileNames, entry.Key)
		}
		if hashNext {
			marker = nextMarker
		} else {
			break
		}
	}
	return
}