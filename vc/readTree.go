package vc

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadTree(treeOid string) {
	err := readTree(treeOid)
	if err != nil {
		log.Fatalf("Error in readTree - %v", err)
	}
}

func readTree(treeOid string) error {
	results, err := getTree(treeOid, "./")
	if err != nil {
		return err
	}

	for path, oid := range results {
		dirName := filepath.Dir(path)
		if _, err := os.Stat(dirName); os.IsNotExist(err) {
			err = os.Mkdir(dirName, 0744)
			if err != nil {
				return err
			}
			fmt.Printf("%v created", dirName)
		}
		err = ioutil.WriteFile(path, []byte(catFile(oid, "")), 0744)
		if err != nil {
			return err
		}
	}
	return nil
}

func getTree(oid, basePath string) (map[string]string, error) {
	results := make(map[string]string)

	for _, i := range iterTreeEntries(oid) {
		type_, oid, name := i[0], i[1], i[2]
		if name == "/" {
			return nil, errors.New("name should not be equal to /")
		}
		if name == ".." || name == "." {
			return nil, errors.New("name should not be equal to .. or .")
		}
		path := filepath.Join(basePath, name)
		if type_ == "blob" {
			results[path] = oid
		} else if type_ == "tree" {
			result, err := getTree(oid, path)
			if err != nil {
				return nil, err
			}
			for k, v := range result {
				results[k] = v
			}
		} else {
			return nil, errors.New("Unknown tree entry")
		}
	}
	return results, nil
}

func iterTreeEntries(oid string) [][]string {
	var treeSlice [][]string

	if oid == "" {
		return nil
	}

	tree := catFile(oid, "tree")
	for _, line := range strings.Split(tree, "\n") {
		hashed := strings.Split(line, " ")
		treeSlice = append(treeSlice, hashed)
	}
	return treeSlice
}
