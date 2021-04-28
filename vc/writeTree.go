package vc

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func WriteTree(path string) {
	oid := writeTree(path)
	fmt.Println(oid)
}

func writeTree(directory string) string {
	oid, type_ := "", ""
	var entries [][]string

	fs, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatalf("Error reading dir [%v] - %v", directory, err)
	}

	for _, f := range fs {
		filePath := filepath.Join(directory, f.Name())
		if f.Name() == VcDir {
			continue
		}
		if f.IsDir() {
			type_ = "tree"
			oid = writeTree(filePath)
		} else {
			type_ = "blob"
			data, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Fatalf("Error reading file [%v] - %v", filePath, err)
			}
			oid = hashObject(data, type_)
		}
		entries = append(entries, []string{type_, oid, filepath.Base(filePath)})
	}

	var hashData []string
	for _, i := range entries {
		s := strings.Join(i, " ")
		hashData = append(hashData, s)
	}
	tree := strings.Join(hashData, "\n")
	return hashObject([]byte(tree), "tree")
}
