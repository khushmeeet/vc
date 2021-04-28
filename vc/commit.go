package vc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func Commit(message string) {
	fmt.Println(commit(message))
}

func commit(message string) string {
	var commit bytes.Buffer
	commit.WriteString("tree ")
	commit.WriteString(writeTree("."))
	commit.WriteString("\n")
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
