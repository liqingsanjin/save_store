package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"save_store/pkg/store"
)

var st store.Store

func main() {
	source := os.Args[1]
	action := os.Args[2]
	st = store.NewQiniu()
	if source == "ceph" {
		st = store.NewCeph()
	}

	switch action {
	case "upload":
		filePath := os.Args[3]
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		path := os.Args[4]
		err = st.Upload(path, fileBytes)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("upload success")
	case "get":
		path := os.Args[3]
		fileBytes, err := st.Get(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = ioutil.WriteFile(os.Args[4], fileBytes, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "list":
		limit, err := strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Println(err)
			return
		}
		list, err := st.List(os.Args[3], limit)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, l := range list {
			fmt.Println(l)
		}
	case "delete":
		err := st.Delete(os.Args[3])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("delete success")
	}
}
