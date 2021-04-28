package vc

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Commit(message string) {
	fmt.Println(commit(message))
}

func commit(message string) string {
	var commit bytes.Buffer

	commit.WriteString("tree ")
	commit.WriteString(writeTree("."))
	commit.WriteString("\n")
	if head, err := getHEAD(); err == nil {
		commit.WriteString("parent ")
		commit.WriteString(head)
		commit.WriteString("\n")
	}
	commit.WriteString("\n")
	commit.WriteString(message)

	oid := hashObject(commit.Bytes(), "commit")
	err := setHEAD(oid)
	if err != nil {
		log.Fatalf("Error setting HEAD - %v", err)
	}

	return oid
}

func setHEAD(oid string) error {
	err := ioutil.WriteFile(filepath.Join(VcDir, "HEAD"), []byte(oid), 0744)
	if err != nil {
		return err
	}

	return nil
}

func getHEAD() (string, error) {
	headPath := filepath.Join(VcDir, "HEAD")
	if _, err := os.Stat(headPath); !os.IsNotExist(err) {
		f, err := ioutil.ReadFile(headPath)
		if err != nil {
			return "", err
		}
		head := strings.Trim(string(f), " ")
		return head, nil
	} else {
		return "", errors.New("HEAD does not exists")
	}
}
