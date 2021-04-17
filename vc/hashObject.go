package vc

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func HashObject(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file [%v] - %v", path, err)
	}
	fmt.Println(hashObject(data, "blob"))
}

func hashObject(data []byte, type_ string) string {
	obj := append([]byte(type_), []byte("\x00")...)
	obj = append(obj, data...)
	hashFunc := sha1.New()
	hashFunc.Write(obj)
	oid := hashFunc.Sum(nil)
	oidHex := fmt.Sprintf("%x", oid)
	filePath := filepath.Join(VcDir, "objects", oidHex)

	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating file [%v] - %v", filePath, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := f.Write(obj); err != nil {
		panic(err)
	}
	return oidHex
}
