package vc

import (
	"log"
	"path/filepath"
)

func Tag(name, oid string) {
	tag := ""
	if oid != "" {
		tag = oid
	} else {
		tag, _ = getRef("HEAD")
	}

	createTag(name, tag)
}

func createTag(name, oid string) {
	tagPath := filepath.Join("tags", name)
	err := updateRef(tagPath, oid)
	if err != nil {
		log.Fatalf("error creating tag - %v", err)
	}
}
