package main

import (
	"apicore"
	_ "apicore/get"
)

func main() {
	err := apicore.Run(":3000")
	if err != nil {
		panic(err)
	}
}
