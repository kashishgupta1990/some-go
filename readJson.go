package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
)

type Config struct {
	Port string `json:"port"`
}

func main(){
	var config *Config
	raw, err := ioutil.ReadFile("./config.json")
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(1);
	}else{
		json.Unmarshal(raw, &config)

		fmt.Println(config.Port)
	}
}
