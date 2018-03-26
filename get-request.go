package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type User struct {
	FirstName string `json:"first_name"`
	Id int `json:"id"`
	LastName string `json:"last_name"`
}

type UserResponse struct {
	Data []User `json:"data"`
	Page int `json:"page"`
}

func makeRequest(url string, c chan UserResponse) {
	var userResponse UserResponse
	resp, _ := http.Get(url)
	bytesBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytesBody, &userResponse)
	c <- userResponse
}

func main() {
	c:= make(chan UserResponse)
	url1:="https://reqres.in/api/users?page=1"
	url2:="https://reqres.in/api/users?page=2"
	url3:="https://reqres.in/api/users?page=3"
	url4:="https://reqres.in/api/users?page=4"
	go makeRequest(url1, c)
	go makeRequest(url2, c)
	go makeRequest(url3, c)
	go makeRequest(url4, c)
	defer fmt.Println(<-c, <-c, <-c, <-c)
}
