package vc

import (
	"fmt"
	"log"
	"path/filepath"
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
