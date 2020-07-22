package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func size(cnt int64) string {
	if cnt == 0 {
		return "empty"
	}
	return strconv.FormatInt(cnt, 10) + "b"
}

func trueDirTree(out io.Writer, path string, printFiles bool, level string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	if !printFiles {
		var newFiles []os.FileInfo
		for _, file := range files {
			if file.IsDir() {
				newFiles = append(newFiles, file)
			}
		}
		files = newFiles
	}
	for idx, file := range files {
		if idx != len(files)-1 {
			if file.IsDir() {
				fmt.Fprintf(out, "%s├───%s\n", level, file.Name())
				trueDirTree(out, filepath.Join(path, file.Name()), printFiles, level+"│\t")
			} else if printFiles {
				fmt.Fprintf(out, "%s├───%s (%s)\n", level, file.Name(), size(file.Size()))
			}
		} else {
			if file.IsDir() {
				fmt.Fprintf(out, "%s└───%s\n", level, file.Name())
				trueDirTree(out, filepath.Join(path, file.Name()), printFiles, level+"\t")
			} else if printFiles {
				fmt.Fprintf(out, "%s└───%s (%s)\n", level, file.Name(), size(file.Size()))
			}
		}
	}
	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	return trueDirTree(out, path, printFiles, "")
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
