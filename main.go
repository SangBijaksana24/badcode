package main

import (
	"fmt"
)

func main() {
	item, err := goodCode("something")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(item)
}
