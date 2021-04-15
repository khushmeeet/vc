package vc

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func CatFile(hash string) {
	content := catFile(hash)
	fmt.Println(content)
}

func catFile(hash string) string {
	hashFilePath := filepath.Join(VcDir, "objects", hash)
	f, err := ioutil.ReadFile(hashFilePath)
	if err != nil {
		log.Fatalf("Error reading hash file [%v] - %v", hashFilePath, err)
	}
	return string(f)
}
