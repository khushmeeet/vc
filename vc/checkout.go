package vc

import (
	"fmt"
	"log"
)

func Checkout(oid string) {
	checkout(oid)
}

func checkout(name string) {
	oid := GetOid(name)
	var head RefValue

	commit, err := getCommit(oid)
	if err != nil {
		log.Fatalf("error getting commit - %v", err)
	}

	err = readTree(commit.Tree)
	if err != nil {
		log.Fatalf("error reading tree commit - %v", err)
	}

	if isBranch(name) {
		head = RefValue{symbolic: true, value: fmt.Sprintf("refs/heads/%v", name)}
	} else {
		head = RefValue{symbolic: false, value: oid}
	}

	err = updateRef("HEAD", head, false)
	if err != nil {
		log.Fatalf("error setting HEAD - %v", err)
	}
}

func isBranch(branch string) bool {
	return getRef(fmt.Sprintf("refs/heads/%v", branch), true).value != ""
}
