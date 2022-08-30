package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
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

func getFileName(name string) string {
	res := filepath.Base(name)
	fmt.Println(res)
	return res
}

func getFileName2(name string) string {
	res := strings.Split(name, "/")
	fmt.Println("getFileName: " + res[len(res)-1])
	return res[len(res)-1]
}

func getInfo(info string, data *exif.Exif) interface{} {
	jdata, err := data.MarshalJSON()
	if err != nil {
		panic(err)
	}
	var jsonObj interface{}
	_ = json.Unmarshal(jdata, &jsonObj)

	var d = jsonObj.(map[string]interface{})
	fmt.Println(d[info])
	return d[info]
}

func setLabel(names ...string) []string {
	var photo []string
	photo = append(photo, "ImageName")
	for _, name := range names {
		photo = append(photo, name)
	}
	return photo
}

func SetLocate(name string) []string {
	file, err := os.Open(name)
	var photo []string
	if err != nil {
		panic(err)
	}

	x, err := exif.Decode(file)
	if err != nil {
		photo = ([]string{getFileName(name), "0", "0"})
		return photo
	}
	lat, lng, err := x.LatLong()
	if err != nil {
		photo = ([]string{getFileName(name), "0", "0"})
		return photo
	}
	photo = ([]string{getFileName(name), strconv.FormatFloat(lat, 'f', 12, 64), strconv.FormatFloat(lng, 'f', 12, 64)})
	fmt.Println(photo)
	return photo
}

func setLatLng(name string, flag *bool) [][]string {
	file, err := os.Open(name)
	var photo [][]string
	if *flag == false {
		photo = append(photo, []string{"ImageName", "lat", "lng"})
	}
	if err != nil {
		panic(err)
	}

	x, err := exif.Decode(file)
	if err != nil {
		photo = append(photo, []string{getFileName(name), "0", "0"})
		return photo
	}
	lat, lng, err := x.LatLong()
	if err != nil {
		photo = append(photo, []string{getFileName(name), "0", "0"})
		return photo
	}
	photo = append(photo, []string{getFileName(name), strconv.FormatFloat(lat, 'f', 12, 64), strconv.FormatFloat(lng, 'f', 12, 64)})
	fmt.Println(photo)
	return photo
}

func setSelectData(name string, info ...string) []string {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	var photo []string
	exif, err := exif.Decode(file)
	if err != nil {
		photo = append(photo, getFileName(name))
		for _, in := range info {
			str := "Not Read" + in
			photo = append(photo, str)
		}
		fmt.Println("sdfklk   ")
		return photo

	} else {
		photo = append(photo, getFileName(name))
		for _, in := range info {
			res := getInfo(in, exif)
			//interface型をstring型に治す
			str, ok := res.(string)
			if !ok {
				str = "fatal type convert"
			}
			photo = append(photo, str)
		}
		return photo
	}

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
	data := dirwalk(path)

	var name string
	fmt.Println("File Name??")
	fmt.Scan(&name)
	f, err := os.Create(name + ".csv")
	if err != nil {
		panic(err)
	}

	var photos [][]string

	var seting string
	fmt.Println("default or custom??")
	fmt.Scan(&seting)
	if seting != "custom" {
		photos = append(photos, setLabel("lat", "lng"))
		for _, d := range data {
			if !strings.Contains(d, "._") {
				photo := SetLocate(d)
				photos = append(photos, photo)
			}
		}
	} else {
		fmt.Println("必要なパラメータ数を入力してください")
		var n int
		fmt.Scan(&n)
		var labels []string
		for i := 0; i < n; i++ {
			var a string
			fmt.Scan(&a)
			labels = append(labels, a)
		}
		photos = append(photos, setLabel(labels...))
		for _, d := range data {
			if !strings.Contains(d, "._") {
				photo := setSelectData(d, labels...)
				photos = append(photos, photo)
			}
		}
	}

	w := csv.NewWriter(f)

	for _, photo := range photos {
		if err := w.Write(photo); err != nil {
			panic(err)
		}
	}
	//バッファに残ってるデータを書き込む
	w.Flush()

}
