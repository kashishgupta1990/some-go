package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}

