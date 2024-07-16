package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool)

	people := [2]string {"nick", "test"}

	for _,person := range people {
		go isCool(person, channel)
		
	}

	fmt.Println(<- channel)
	fmt.Println(<- channel)
}

func isCool(person string, channel chan bool)  {
	fmt.Println("isCool request received")
	time.Sleep(time.Second * 2)

	fmt.Println(person, "is Cool")

	channel <- true
}