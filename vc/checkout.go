package vc

import "log"

func Checkout(oid string) {
	checkout(oid)
}

func checkout(oid string) {
	commit, err := getCommit(oid)
	if err != nil {
		log.Fatalf("error getting commit - %v", err)
	}

	err = readTree(commit.Tree)
	if err != nil {
		log.Fatalf("error reading tree commit - %v", err)
	}

	err = updateRef("HEAD", oid)
	if err != nil {
		log.Fatalf("error setting HEAD - %v", err)
	}
}
