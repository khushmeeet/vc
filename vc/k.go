package vc

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func K() {
	var oids []string
	var dot bytes.Buffer
	dot.WriteString("digraph commits {\n")

	refs := iterRefs()
	refGen := yieldRefs(refs)
	for _ = range refs {
		refName, ref, err := refGen()
		if err != nil {
			panic(err)
		}
		dot.WriteString(fmt.Sprintf("\"%v\" [shape=note]\n", refName))
		dot.WriteString(fmt.Sprintf("\"%v\" -> \"%v\"\n", refName, ref.value))
		oids = append(oids, ref.value)
	}

	for _, oid := range iterCommitsAndParents(oids...) {
		commit, err := getCommit(oid)
		if err != nil {
			log.Fatalf("error getting commit - %v", err)
		}
		dot.WriteString(fmt.Sprintf("\"%v\" [shape=box style=filled label=\"%v\"]\n", oid, oid[:10]))
		if commit.Parent != "" {
			dot.WriteString(fmt.Sprintf("\"%v\" -> \"%v\"\n", oid, commit.Parent))
		}
	}

	dot.WriteString("}")
	fmt.Println(dot.String())

	dotCmd := exec.Command("dot", "-Tpng", "/dev/stdin")
	dotCmdIn, err := dotCmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	dotCmdOut, _ := dotCmd.StdoutPipe()
	err = dotCmd.Start()
	if err != nil {
		panic(err)
	}
	_, err = dotCmdIn.Write(dot.Bytes())
	if err != nil {
		panic(err)
	}
	err = dotCmdIn.Close()
	if err != nil {
		panic(err)
	}
	b, _ := ioutil.ReadAll(dotCmdOut)
	ioutil.WriteFile("vc.png", b, 0755)
	err = dotCmd.Wait()
	if err != nil {
		panic(err.Error())
	}

	//dotCmdOut.Read()
}

func iterCommitsAndParents(oids ...string) []string {
	oidsList := oids
	var visited []string
	var iters []string

	for i := 0; i < len(oidsList); i++ {
		oid := oidsList[len(oidsList)-1]
		if oid == "" || contains(visited, oid) {
			continue
		}
		visited = append(visited, oid)
		iters = append(iters, oid)
		commit, err := getCommit(oid)
		if err != nil {
			log.Fatalf("error getting commit - %v", err)
		}
		oidsList = append(oidsList, commit.Parent)
	}
	return iters
}

func iterRefs() []string {
	refs := []string{"HEAD"}

	err := filepath.WalkDir(filepath.Join(VcDir, "refs"), func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			relPath := strings.TrimPrefix(path, VcDir)
			refs = append(refs, relPath)
			return nil
		}
		return nil
	})
	if err != nil {
		log.Fatalf("error walking .vc -%v", err)
	}

	return refs
}

func yieldRefs(refs []string) func() (string, RefValue, error) {
	refsLen := len(refs)
	n := 0
	return func() (string, RefValue, error) {
		if n < refsLen {
			oid, err := getRef(refs[n])
			if err != nil {
				log.Fatalf("%v", err)
			}
			ref := refs[n]
			n = n + 1
			return ref, oid, nil
		}
		return "", RefValue{}, errors.New("yield complete")
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
