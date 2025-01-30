package main

import (
	"fmt"
	"github.com/jasonlovesdoggo/gopen"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 || len(argsWithoutProg) > 1 {
		fmt.Println("usage: gopen <URL>")
		os.Exit(1)
	}

	// open the URL in the default browser
	err := gopen.Open(argsWithoutProg[0])
	if err != nil {
		fmt.Println(err)
	}
}
