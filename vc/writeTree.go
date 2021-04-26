package vc

import (
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func WriteTree(path string) {
	writeTree(path)
}

func writeTree(directory string) string {
	oid, type_ := "", ""
	var entries [][]string
	err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && filepath.Base(path) == VcDir {
			return filepath.SkipDir
		}
		if d.IsDir() {
			type_ = "tree"
			oid = writeTree(path)
		}
		if !d.IsDir() {
			type_ = "blob"
			f, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			oid = hashObject(f, type_)
		}
		entries = append(entries, []string{filepath.Base(path), oid, type_})
		return nil
	})
	if err != nil {
		log.Fatalf("Error scanning dir [%v] - %v", directory, err)
	}

	var hashData []string
	for _, i := range entries {
		s := strings.Join(i, " ")
		hashData = append(hashData, s)
	}
	tree := strings.Join(hashData, "\n")
	return hashObject([]byte(tree), "tree")
}
