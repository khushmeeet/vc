package vc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func DoCommit(message string) {
	fmt.Println(commit(message))
}

func commit(message string) string {
	var commit bytes.Buffer
	head := getRef("HEAD", true)

	commit.WriteString("tree ")
	commit.WriteString(writeTree("."))
	commit.WriteString("\n")

	if head.value != "" {
		commit.WriteString("parent ")
		commit.WriteString(head.value)
		commit.WriteString("\n")
	}

	commit.WriteString("\n")
	commit.WriteString(message)

	oid := hashObject(commit.Bytes(), "commit")
	err := updateRef("HEAD", RefValue{symbolic: false, value: oid}, true)
	if err != nil {
		log.Fatalf("Error setting HEAD - %v", err)
	}

	return oid
}

func updateRef(ref string, rv RefValue, deref bool) error {
	value := ""
	r, _ := getRefInternal(ref, deref)

	if rv.symbolic {
		value = fmt.Sprintf("ref: %v", rv.value)
	} else {
		value = rv.value
	}

	refPath := filepath.Join(VcDir, r)
	err := os.MkdirAll(strings.TrimSuffix(refPath, filepath.Base(refPath)), 0766)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(refPath, []byte(value), 0766)
	if err != nil {
		return err
	}

	return nil
}

func getRef(ref string, deref bool) RefValue {
	ref, val := getRefInternal(ref, deref)
	return val
}

func getRefInternal(ref string, deref bool) (string, RefValue) {
	value := ""

	refPath := filepath.Join(VcDir, ref)
	if stat, err := os.Stat(refPath); !os.IsNotExist(err) {
		if !stat.IsDir() {
			f, err := ioutil.ReadFile(refPath)
			if err != nil {
				return "", RefValue{}
			}
			value = strings.Trim(string(f), " ")
		}
	}

	symbolic := value != "" && strings.HasSuffix(value, "ref:")
	if symbolic {
		refV := strings.Trim(strings.Split(value, ":")[1], " ")
		if deref {
			return getRefInternal(refV, true)
		}
	}
	return ref, RefValue{symbolic: symbolic, value: value}
}
