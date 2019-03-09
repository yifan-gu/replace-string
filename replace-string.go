package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	usage = "Usage: replace-string ROOT old_string new_string"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println(usage)
		os.Exit(1)
	}

	root, old, new := os.Args[1], os.Args[2], os.Args[3]

	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		if !bytes.Contains(data, []byte(old)) {
			return nil
		}

		fmt.Println("Update", path)

		newData := bytes.ReplaceAll(data, []byte(old), []byte(new))
		return ioutil.WriteFile(path, newData, info.Mode())
	}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
