package vc

import (
	"fmt"
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

func GetOid(name string) string {
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

	//ref, err := getRef(name)
	//if err != nil {
	//	log.Fatalf("error getting ref - %v", err)
	//}
	//
	//if ref != "" {
	//	return ref
	//} else {
	//	return name
	//}
}

func createTag(name, oid string) {
	tagPath := filepath.Join("refs", "tags", name)
	err := updateRef(tagPath, oid)
	if err != nil {
		log.Fatalf("error creating tag - %v", err)
	}
}
