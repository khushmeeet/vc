package vc

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func K() {
	var oids []string

	refs := iterRefs()
	refGen := yieldRefs(refs)
	for _ = range refs {
		refName, ref, err := refGen()
		if err != nil {
			panic(err)
		}
		fmt.Println(refName, ref)
		oids = append(oids, ref)
	}

	for _, oid := range iterCommitsAndParents(oids...) {
		commit, err := getCommit(oid)
		if err != nil {
			log.Fatalf("error getting commit - %v", err)
		}
		fmt.Println(oid)
		if commit.Parent != "" {
			fmt.Println("Parent ", commit.Parent)
		}
	}
}

func iterCommitsAndParents(oids ...string) []string {
	var visited []string
	var iters []string

	for _, oid := range oids {
		if oid == "" || contains(visited, oid) {
			continue
		}
		visited = append(visited, oid)
		iters = append(iters, oid)
		commit, err := getCommit(oid)
		if err != nil {
			log.Fatalf("error getting commit - %v", err)
		}
		oids = append(oids, commit.Parent)
	}
	return iters
}

func iterRefs() []string {
	refs := []string{"HEAD"}

	err := filepath.WalkDir(filepath.Join(VcDir, "refs"), func(path string, d fs.DirEntry, err error) error {
		refs = append(refs, path)
		return nil
	})
	if err != nil {
		log.Fatalf("error walking .vc -%v", err)
	}

	return refs
}

func yieldRefs(refs []string) func() (string, string, error) {
	refsLen := len(refs)
	n := 0
	return func() (string, string, error) {
		if n < refsLen {
			oid, _ := getRef(refs[n])
			ref := refs[n]
			n = n + 1
			return ref, oid, nil
		}
		return "", "", errors.New("yield complete")
	}
}

func contains(arr []string, v string) bool {
	for _, i := range arr {
		if i == v {
			return true
		}
	}
	return false
}
