package main

import (
	"fmt"

	"github.com/jaewon1022/learngo/src/github.com/jaewon1022/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first":"First Word"}
	key := "second"
	value := "Second Word"
	newValue := "fdklsjhfhlsakd"


	err := dictionary.Add(key, value)
	if(err != nil) {
		fmt.Println(err)
		} else {
		updateErr := dictionary.Update(key, newValue)

		if(updateErr != nil) {
			fmt.Println(updateErr)
		} else {
			dict := dictionary.Info()
	
			fmt.Println("dict list : ", dict)

			dictionary.Delete(key)

			afterDict := dictionary.Info()

			fmt.Println("after delete : ", afterDict)
		}
	}
}
