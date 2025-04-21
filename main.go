package main

import (
	// "fmt"
	functions "myls1/functions"
	"os"
)

func main() {
	args := os.Args
	Themap, _ := functions.ArgumentChecking(args)
	// fmt.Println(Themap)
	// fmt.Println(files)

	if Themap["recursive"]{
		Themap["R"] = true
	}
	if Themap["time"]{
		Themap["t"] = true
	}
	if Themap["all"]{
		Themap["ls"] = true
		Themap["l"] = true
		Themap["R"] = true
		Themap["r"] = true
		Themap["a"] = true
		Themap["t"] = true
	}
	if Themap["reverse"]{
		Themap["r"] = true
	}

	// for _, x := range files{
	// 	if Themap
	// }
}

