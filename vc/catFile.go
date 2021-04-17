package vc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func CatFile(hash string) {
	content := catFile(hash, "")
	fmt.Println(content)
}

func catFile(oid, expected string) string {
	hashFilePath := filepath.Join(VcDir, "objects", oid)
	f, err := ioutil.ReadFile(hashFilePath)
	if err != nil {
		log.Fatalf("Error reading oid file [%v] - %v", hashFilePath, err)
	}
	content := bytes.Split(f, []byte("\x00"))
	type_, data := content[0], content[1]
	if expected != "" && string(type_) != expected {
		panic("Object type mismatch!")
	}
	return string(data)
}
