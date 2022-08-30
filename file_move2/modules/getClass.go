package modules

import (
	"encoding/json"
	"io/ioutil"
)

type Koke interface {
	ImageClass()
}

//jsonを受け取るための構造体
type Image struct {
	ImageName string `json:"Imagename"`
	Number    string `json:"class"`
}

func Image_Class(name string) ([]Image, error) {
	byte, err := ioutil.ReadFile(name)

	if err != nil {
		panic(err)
	}

	var images []Image
	err = json.Unmarshal(byte, &images)

	if err != nil {
		panic(err)
	}

	return images, err
}
