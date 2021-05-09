package vc

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func DoCommit(message string) {
	fmt.Println(commit(message))
}

func commit(message string) string {
	var commit bytes.Buffer

	commit.WriteString("tree ")
	commit.WriteString(writeTree("."))
	commit.WriteString("\n")
	if head, err := getRef("HEAD"); err == nil {
		commit.WriteString("parent ")
		commit.WriteString(head.value)
		commit.WriteString("\n")
	}
	commit.WriteString("\n")
	commit.WriteString(message)

	oid := hashObject(commit.Bytes(), "commit")
	err := updateRef("HEAD", RefValue{symbolic: false, value: oid})
	if err != nil {
		log.Fatalf("Error setting HEAD - %v", err)
	}

	return oid
}

func updateRef(ref string, rv RefValue) error {
	refPath := filepath.Join(VcDir, ref)
	err := os.MkdirAll(strings.TrimSuffix(refPath, filepath.Base(refPath)), 0766)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(refPath, []byte(rv.value), 0766)
	if err != nil {
		return err
	}

	return nil
}

func getRef(ref string) (RefValue, error) {
	refPath := filepath.Join(VcDir, ref)
	if stat, err := os.Stat(refPath); !os.IsNotExist(err) {
		if !stat.IsDir() {
			f, err := ioutil.ReadFile(refPath)
			if err != nil {
				return RefValue{}, err
			}
			value := strings.Trim(string(f), " ")
			if strings.HasSuffix(value, "ref:") {
				return getRef(strings.TrimSuffix(strings.Split(value, ":")[1], " "))
			}
			return RefValue{symbolic: false, value: value}, nil
		}
		return RefValue{}, errors.New(fmt.Sprintf("%v is not a file", refPath))
	} else {
		return RefValue{}, errors.New(fmt.Sprintf("%v does not exists", refPath))
	}
}
