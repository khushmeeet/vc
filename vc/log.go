package vc

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type Commit struct {
	Tree    string
	Parent  string
	Message string
}

func Log(oid string) {
	head := oid

	for head != "" {
		commit, err := getCommit(head)
		if err != nil {
			log.Fatalf("error getting commit - %v", err)
		}

		fmt.Printf("commit %v\n", head)
		fmt.Printf("%v\n", commit.Message)
		fmt.Println("------------------------")
		head = commit.Parent
	}
}

func getCommit(oid string) (Commit, error) {
	commit := catFile(oid, "commit")
	lines := strings.Split(commit, "\n")
	treeHash, parentHash := "", ""

	for _, line := range lines {
		if line != "" {
			hashObj := strings.Split(line, " ")
			key, value := hashObj[0], hashObj[1]
			if key == "tree" {
				treeHash = value
			} else if key == "parent" {
				parentHash = value
			} else {
				return Commit{}, errors.New(fmt.Sprintf("unknown key %v", key))
			}
		} else {
			break
		}
	}
	message := strings.Join(lines, "\n")
	return Commit{Tree: treeHash, Parent: parentHash, Message: message}, nil
}
