package main

import (
	"fmt"
	"time"
)

func sum(list []int, c chan int, delay time.Duration, quitChan chan bool, quit bool) {
	ans := 0
	for _, i := range list {
		ans += i
	}

	time.Sleep(delay * time.Second)
	//fmt.Println("Yo: ",  <- c)
	c <- ans

	if quit{
		quitChan <- true
	}
}

func main() {
	arr1 := []int{2, 3}
	arr2 := []int{4, 9}
	arr3 := []int{2}
	c := make(chan int, 1)
	quit := make(chan bool)

	go sum(arr1, c, 5, quit, false)
	go sum(arr2, c, 7, quit, true)
	go sum(arr3, c, 1, quit, false)

	for{
		select {
		case x:= <- c:
			fmt.Println("I am ", x)
		case <-quit:
			fmt.Println("I am quiting.")
			return
		default:
			fmt.Print(".")
			time.Sleep(1000*time.Microsecond)
		}
	}

}
