package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func readcsv(path string) [][]string {
	var res [][]string
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		println(path)
		panic(err)
	}

	for _, v := range rows {
		res = append(res, v)
	}
	return res
}

func readLabel(data [][]string, usetype bool) (string, string) {
	label := ""
	quetion := "?" //インジェクション対策のための文字列　なんか他に方法ないんか
	create := "id INTEGER PRIMARY KEY AUTOINCREMENT"
	for i := 0; i < len(data[0]); i++ {
		if data[0][i] == "index" {
			data[0][i] = data[0][i] + "2"
		}
		//if i != 0 {quetion += ",?"}
		label += ("," + "'" + data[0][i] + "'")
	}
	if usetype {
		return create + label, quetion
	} else {
		return label[1:], quetion
	}
}

func data_shaping(data []string) string {
	var res string
	res = data[0]
	for i := 1; i < len(data); i++ {
		//res += ("," + data[i])
		if data[i] == "" {
			data[i] = "dummy"
		}
		res += ("," + "'" + data[i] + "'")
	}
	return res
}

func dbConnect(name string) (*sql.DB, error) {
	DB, err := sql.Open("sqlite3", "./"+name)
	return DB, err
}

func create_table(filename, name, label string) {
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
		)
	`, name, label)
	DB, err := dbConnect(filename)
	defer DB.Close()
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(cmd)
	if err != nil {
		panic(err)
	}
}

func insert_data(filename, tablename, label, data, q string) {
	DB, err := dbConnect(filename)
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	cmd := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, tablename, label, data)
	fmt.Println(cmd)
	res, err := DB.Exec(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	/*stmt, err := DB.Prepare(cmd)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()*/

}

func main() {
	var input [3]string
	fmt.Println("filename input_db table")
	fmt.Scan(&input[0], &input[1], &input[2])
	data := readcsv(input[0])
	fmt.Println(data[1])

	res := data_shaping(data[1])
	fmt.Println(res)
	fmt.Println(readLabel(data, true))
	label, _ := readLabel(data, true)
	fmt.Println(label)
	//create_table("b.db", "ful", label)
	create_table(input[1], input[2], label)
	insert_label, q := readLabel(data, false)
	for _, d := range data[1:] {
		//insert_data("b.db", "ful", insert_label, data_shaping(d), q)
		insert_data(input[1], input[2], insert_label, data_shaping(d), q)
	}
	//insert_data("b.db", "ful", insert_label, data_shaping(data[2]), q)

}
