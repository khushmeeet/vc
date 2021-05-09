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
		commit.WriteString(head)
		commit.WriteString("\n")
	}
	commit.WriteString("\n")
	commit.WriteString(message)

	oid := hashObject(commit.Bytes(), "commit")
	err := updateRef("HEAD", oid)
	if err != nil {
		log.Fatalf("Error setting HEAD - %v", err)
	}

	return oid
}

func updateRef(ref, oid string) error {
	refPath := filepath.Join(VcDir, ref)
	err := os.MkdirAll(strings.TrimSuffix(refPath, filepath.Base(refPath)), 0766)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(refPath, []byte(oid), 0766)
	if err != nil {
		return err
	}

	return nil
}

func getRef(ref string) (string, error) {
	refPath := filepath.Join(VcDir, ref)
	if stat, err := os.Stat(refPath); !os.IsNotExist(err) {
		if !stat.IsDir() {
			f, err := ioutil.ReadFile(refPath)
			if err != nil {
				return "", err
			}
			value := strings.Trim(string(f), " ")
			if strings.HasSuffix(value, "ref:") {
				parentRef, err := getRef(strings.TrimSuffix(strings.Split(value, ":")[1], " "))
				return parentRef, err
			}
			return value, nil
		}
		return "", errors.New(fmt.Sprintf("%v is not a file", refPath))
	} else {
		return "", errors.New(fmt.Sprintf("%v does not exists", refPath))
	}
}
