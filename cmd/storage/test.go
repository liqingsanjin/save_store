package main

import (
	"save_store/pkg/store"
	"fmt"
)

func main()  {
	qiniu := store.NewStorage("liqingsanjin", "http://oxx1uk7mm.bkt.clouddn.com",
		"VjgieAVwG723rZomW6h0SRdEIXHe_vNnBDtDw527", "bnQc-7eEPnnN5LEolwfxbAQjwr07d85CShh9lfLg",
		"qiniu", false)
	testUpload(qiniu)
	testGet(qiniu)
	testList(qiniu)
	testDelete(qiniu)

	ceph := store.NewStorage("testceph", "cs54:7080",
		"Z8DFXNIGF71P8K71D0ZA", "VI0k5TecQOjwowWzhfapKRUfJfhs041hphieEhIM",
		"ceph", false)
	testUpload(ceph)
	testGet(ceph)
	testList(ceph)
	testDelete(ceph)
}

func testUpload(st store.Storage) {
	err := st.Upload("abc.txt", []byte("hello, world"))
	if err != nil {
		fmt.Println(err)
	}
}

func testGet(st store.Storage) {
	fileBytes, err := st.Get("abc.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(fileBytes))
}

func testList(st store.Storage) {
	fileNames, err := st.List("", "",10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileNames)
}

func testDelete(st store.Storage) {
	err := st.Delete("abc.txt")
	if err != nil {
		fmt.Println(err)
	}
}