package vc

import "fmt"

func Status() {
	head := GetOid("@")
	branch := getBranchName()
	if branch != "" {
		fmt.Println(fmt.Sprintf("On Branch -> %v", branch))
	} else {
		fmt.Println(fmt.Sprintf("HEAD detached at %v", head[:10]))
	}
}
