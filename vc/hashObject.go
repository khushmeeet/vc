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
	fmt.Println(hashObject(data))
}

func hashObject(data []byte) string {
	hashFunc := sha1.New()
	hashFunc.Write(data)
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

	if _, err := f.Write(data); err != nil {
		panic(err)
	}
	return oidHex
}
