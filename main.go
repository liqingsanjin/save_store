package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"save_store/store"
)

var st store.Store

func main() {
	//st = store.NewQiniu()
	st = store.NewCeph()
	file, err := os.Open("/Users/admin/aaa.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	path := "img/aaa.png"
	err = st.Upload(path, fileBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileBytes, err = st.Get(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fileBytes)
	// fileNames, err := st.List("img/", 10)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for _, fileName := range fileNames {
	// 	fmt.Println(fileName)
	// }
	// err = st.Delete(path)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
