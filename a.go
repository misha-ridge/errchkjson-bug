package main

import (
	"encoding/json"
	"fmt"
)

func check(a interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return a
}

func main() {
	fmt.Println(check(json.Marshal(47)))
}
