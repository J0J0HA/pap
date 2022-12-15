package main

import (
	"io/fs"
	"os"
)

func WriteFile(name string, text string, perms fs.FileMode) {
	err := os.WriteFile(name, []byte(text), perms)
	Error(err, "an error occurred while writing %s", name)
}

func MakeDirectory(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	Error(err, "an error occurred while creating %s", path)
}
