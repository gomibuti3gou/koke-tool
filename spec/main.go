package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths
}

//
func readCsv(path string, rownum int) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	var res []string
	for _, v := range rows {
		fmt.Println(v[rownum])
		res = append(res, v[rownum])
	}
	return res
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

	var name string
	fmt.Println("File Name??")
	fmt.Scan(&name)
	outputFile, err := os.Create(name + ".csv")
	if err != nil {
		panic(err)
	}

	var fullData [][]string
	paths := dirwalk(path)
	for i := 0; i < len(paths); i++ {
		if i == 0 {
			fullData = append(fullData, readCsv(paths[i], 0))
		}
		fullData = append(fullData, readCsv(paths[i], 1))
	}

	w := csv.NewWriter(outputFile)

	for _, data := range fullData {
		if err := w.Write(data); err != nil {
			panic(err)
		}
	}
	//バッファに残ってるデータを書き込む
	w.Flush()

}
