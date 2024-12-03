package main

import (
	"fmt"
	"os"
)

func main() {
	dir, _ := getCurrentDir()
	fmt.Println(dir)
}

func getCurrentDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}
