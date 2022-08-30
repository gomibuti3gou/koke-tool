package modules

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func DirSearch(dir string) []string {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		//ディレクトリかを判定。ディレクトリなら再帰的に探索
		if file.IsDir() {
			paths = append(paths, DirSearch(filepath.Join(dir, file.Name()))...)
			continue
		}
		//カレントディレクトリとファイル名を結合
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func DirSearchFull(dir string) []string {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		//ディレクトリかを判定。ディレクトリなら再帰的に探索
		if file.IsDir() {
			paths = append(paths, DirSearch(filepath.Join(dir, file.Name()))...)
			continue
		}
		//カレントディレクトリとファイル名を結合
		absfilepath, err := filepath.Abs(filepath.Join(dir, file.Name()))
		if err != nil {
			panic(err)
		}
		paths = append(paths, absfilepath)
	}

	return paths
}

func GetFileName(name string) string {
	res := filepath.Base(name)
	return res
}

//ディレクトリの文字列の中で、ファイル名を抽出
func GetFileName2(name string) string {
	res := strings.Split(name, "/")
	return res[len(res)-1]
}
