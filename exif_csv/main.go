package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
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

func change_type(value interface{}) string {
	res := ""
	switch v := value.(type) {
	case int:
		res = fmt.Sprintf("%d", value)
	case float32:
		res = fmt.Sprintf("%f", value)
	case float64:
		res = fmt.Sprintf("%f", value)
	case string:
		res = fmt.Sprintf("%s", value)
	case []string:
		vv, _ := value.([]string)
		res = strings.Join(vv, " ")
	case []int:
		vv, _ := value.([]string)
		res = strings.Join(vv, " ")
	case []float64:
		vv, _ := value.([]string)
		res = strings.Join(vv, " ")
	case []interface{}:
		for _, v2 := range v {
			v2_str := change_type(v2)
			res += v2_str + ","
		}
		res = res[0 : len(res)-1]
	default:
		res = fmt.Sprint(value)
	}
	return res
}

func exif_info(name string) (map[string]interface{}, error) {
	var datas map[string]interface{}
	file, err := os.Open(name)
	if err != nil {
		return datas, err
	}
	exif, err := exif.Decode(file)
	if err != nil {
		return datas, err
	}
	jdata, err := exif.MarshalJSON()
	if err != nil {
		return datas, err
	}
	var jsonObj interface{}
	_ = json.Unmarshal(jdata, &jsonObj)
	datas = jsonObj.(map[string]interface{})
	fmt.Println(datas)
	return datas, err
}

func getAllLabel(name string) ([]string, error) {
	datas, err := exif_info(name)
	var labels []string
	if err != nil {
		fmt.Println("fdjlsak")
		return labels, err
	}
	for l, _ := range datas {
		labels = append(labels, l)
	}
	sort.Strings(labels)
	labels = append([]string{"ImageName"}, labels...)
	return labels, err
}

func allData(name string, labels []string) ([]string, error) {
	var photo []string
	datas, err := exif_info(name)
	if err != nil {
		return photo, err
	}

	for i := 1; i < len(labels); i++ {
		var str string
		str = change_type(datas[labels[i]])
		photo = append(photo, str)
	}
	photo = append([]string{getFileName(name)}, photo...)

	fmt.Println(photo)
	return photo, err
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
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var path string
	if err != nil {
		panic(err)
	}
	fmt.Printf("Current Folder %s \n", pwd)
	fmt.Println("Input Folder Path ↓↓")
	fmt.Scan(&path)
	data := dirwalk(path)
	//fmt.Println(data)

	var name string
	fmt.Println("File Name??")
	fmt.Scan(&name)
	f, err := os.Create(name + ".csv")
	if err != nil {
		panic(err)
	}

	var photos [][]string
	var seting string
	var flag bool = true
	fmt.Println("default or custom or all ??")
	fmt.Scan(&seting)
	fmt.Println(data[0])

	if seting == "all" {
		for _, d := range data {
			if flag {
				labels, err := getAllLabel(d)
				if err != nil {
					continue
				}
				photos = append(photos, labels)
				flag = false
			}
			photo, err := allData(d, photos[0])
			if err != nil {
				continue
			}
			photos = append(photos, photo)
		}
	} else if seting != "custom" {
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
