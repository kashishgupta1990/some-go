package main

import (
	"os"
	"fmt"
)

func main() {
	arguments:=os.Args

	fmt.Println(arguments[1:])
}
