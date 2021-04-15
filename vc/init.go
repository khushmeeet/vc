package vc

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Init() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current dir - %v", err)
	}

	if err = os.Mkdir(VcDir, 0744); err != nil {
		log.Fatalf("Error creating dir [%v] - %v", VcDir, err)
	}

	if err = os.Mkdir(filepath.Join(VcDir, "objects"), 0744); err != nil {

	}

	fmt.Printf("Initialized empty vc repository in %v/%v", currentDir, VcDir)
}
