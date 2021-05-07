package vc

import (
	"fmt"
	"log"
	"path/filepath"
)

func Tag(name, oid string) {
	createTag(name, oid)
}

func GetOid(name string) string {
	if name == "@" {
		name = "HEAD"
	}

	refsToTry := [4]string{
		fmt.Sprintf("%v", name),
		fmt.Sprintf("refs/%v", name),
		fmt.Sprintf("refs/tags/%v", name),
		fmt.Sprintf("refs/heads/%v", name),
	}

	for _, ref := range refsToTry {
		oid, err := getRef(ref)
		if err == nil {
			return oid
		}
	}

	if len(name) == 40 {
		return name
	} else {
		return ""
	}
}

func createTag(name, oid string) {
	tagPath := filepath.Join("refs", "tags", name)
	err := updateRef(tagPath, oid)
	if err != nil {
		log.Fatalf("error creating tag - %v", err)
	}
}
