package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filepath := "./read-file/hello.json"
	data, err := readAll(filepath)
	if err != nil {
		panic(err)
	}
	users := make([]*user, 0)
	if err := json.Unmarshal(data, &users); err != nil {
		panic(err)
	}
	for _, v := range users {
		fmt.Println(v.Name, " => ", v.Age)
	}
}

type user struct {
	Name string `json:"name"` // 姓名
	Age  int    `json:"age"`  // 年龄
}

func readAll(filepath string) ([]byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
