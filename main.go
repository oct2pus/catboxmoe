package main

import (
	"catboxmoe/api"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		help()
	} else {
		/*file, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer file.Close()*/
		lol, err := api.Upload(os.Args[1])
		if err != nil {
			panic(err)
		}
		println(string(lol))
	}
}

func help() {
	println("lol")
}
