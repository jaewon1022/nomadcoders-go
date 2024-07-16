package main

import (
	"fmt"
	"time"
)

func main() {
	go count("aaa")
	go count("bbb")
	time.Sleep(time.Second * 5)
}

func count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, i)
		time.Sleep(time.Second * 1)
	}
}