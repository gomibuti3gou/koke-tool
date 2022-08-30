package main

import (
	"file_move/modules"
	"fmt"
	"os"
	"path/filepath"
)

func fileMove(dataPath, jsonPath string) error {
	dirPaths := modules.DirSearch(dataPath)
	//fmt.Println(dirPaths)

	images, err := modules.Image_Class(jsonPath)

	if err != nil {
		panic(err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		for _, path := range dirPaths {
			filepaths := filepath.Join(pwd, path)
			dirPath := filepath.Join(pwd, "class", image.Number, filepath.Base(path))
			//ファイルとJsonのImageNameが一致　ちゃんとファイルがあるかも確かめる
			_, err := os.Stat(filepath.Dir(dirPath))
			fmt.Println("path", filepath.Dir(dirPath))
			fmt.Println("image", image)
			if err != nil {
				if err := os.Mkdir(filepath.Dir(dirPath), 0777); err != nil {
					fmt.Println("32", err)
					panic(err)
				}
			}

			_, err = os.Stat(path)
			if modules.GetFileName(path) == image.ImageName && err == nil {
				fmt.Println("filename : " + filepaths)
				fmt.Println("dirname : " + dirPath)
				//ファイル移動
				err := os.Rename(filepaths, dirPath)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	return err
}

func filemoveByPath(imagePath, jsonPath string) error {
	imagePaths := modules.DirSearchFull(imagePath)
	jsonPaths, err := modules.Image_Class(jsonPath)
	fmt.Println(imagePaths)
	fmt.Println(jsonPaths)
	if err != nil {
		panic(err)
	}
	pwd := filepath.Dir(imagePaths[0])
	for _, image := range jsonPaths {
		for _, path := range imagePaths {
			//filepaths := filepath.Join(filepath.Dir(path), path)
			filepaths := path
			dirPath := filepath.Join(pwd, image.Number, filepath.Base(path))

			//ファイルとJsonのImageNameが一致　ちゃんとファイルがあるかも確かめる
			_, err := os.Stat(filepath.Dir(dirPath))
			fmt.Println("path", filepaths)
			fmt.Println("image", dirPath)
			if err != nil {
				if err := os.Mkdir(filepath.Dir(dirPath), 0777); err != nil {
					fmt.Println("32", err)
					panic(err)
				}
			}

			_, err = os.Stat(path)
			if modules.GetFileName(path) == image.ImageName && err == nil {
				fmt.Println("filename : " + filepaths)
				fmt.Println("dirname : " + dirPath)
				//ファイル移動
				err := os.Rename(filepaths, dirPath)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	return err
}

func main() {
	fmt.Println("画像フォルダのパスを入力してください。 ↓↓↓")
	var inputImage, answer string
	fmt.Scan(&inputImage)
	fmt.Println("Jsonファイルのパスを入力してください。↓↓↓")
	fmt.Scan(&answer)
	err := filemoveByPath(inputImage, answer)

	if err != nil {
		panic(err)
	}
}
