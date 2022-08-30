package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func dirpath(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var paths []string
	for _, file := range files {
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	fmt.Println(paths)
	return paths
}

func main() {
	var path string
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Current Folder %s \n", pwd)
	fmt.Println("Input Folder Path ↓↓")
	fmt.Scan(&path)
	paths := dirpath(path)
	for _, p := range paths {
		if err := os.Rename(p, (p)+".csv"); err != nil {
			fmt.Println(err)
		}
	}

}
