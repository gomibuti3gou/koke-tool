package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Json map[string]interface{}

func read_csv(path, output string) error {
	res := []Json{}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)
	//rows, err := r.ReadAll()
	headers, err := r.Read()
	if err != nil {
		return err
	}

	for {
		rows, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		jsondata := make(Json)
		for i := range rows {
			jsondata[headers[i]] = string(rows[i])
		}
		res = append(res, jsondata)
	}

	jsons, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(output, jsons, 0666)
	if err != nil {
		return err
	}
	return err
}

func main() {
	var input [2]string
	fmt.Println("filename output_filename ")
	fmt.Scan(&input[0], &input[1])
	err := read_csv(input[0], input[1])
	if err != nil {
		fmt.Println(err)
	}
}
