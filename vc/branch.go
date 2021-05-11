package vc

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func Branch(name, oid string) {
	err := createBranch(name, oid)
	if err != nil {
		log.Fatalf("error creating branch - %v\n", err)
	}

	fmt.Printf("Branch %v created at %v\n", name, oid)
}

func createBranch(name, oid string) error {
	err := updateRef(filepath.Join("refs", "heads", name), RefValue{symbolic: false, value: oid}, true)
	if err != nil {
		return err
	}

	return nil
}

func getBranchName() string {
	head := getRef("HEAD", false)
	if !head.symbolic {
		return ""
	}
	headVal := head.value
	if strings.HasPrefix(headVal, "refs/heads/") {
		relpath, _ := filepath.Rel("refs/heads/", headVal)
		return relpath
	}
	return ""
}
