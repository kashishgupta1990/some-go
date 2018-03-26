package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error  {
	fileName:= p.Title + ".txt"
	return ioutil.WriteFile(fileName, p.Body, 0600)
}

func (p *Page) readPage(title string) (*Page, error) {
	fileName:=title + ".txt"
	body, err := ioutil.ReadFile(fileName)
	if err != nil{
		return nil, err
	}else{
		return &Page{Title:title, Body:body}, nil
	}
}



func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, req *http.Request){
	fmt.Println(req.URL.Path[1:])
	fmt.Fprintf(w, "Hi, Their")
}
