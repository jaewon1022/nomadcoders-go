package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	people := [3]string {"nick", "test", "bm", }

	for _,person := range people {
		go isCool(person, channel)
		
	}

	for i := 0; i < len(people) ; i++ {
		fmt.Println(<-channel)
	}
}

func isCool(person string, channel chan string)  {
	time.Sleep(time.Second * 2)

	channel <- person + " is Cool"
}