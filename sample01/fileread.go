package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileContent, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Err")
	}

	fmt.Println(string(fileContent))
