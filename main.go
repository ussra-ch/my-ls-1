package main

import (
	"fmt"
	functions "myls1/functions"
)

func main() {
	// args := os.Args
	// Themap, files := argument_check(args)
	// fmt.Println(args)
	// fmt.Println(Themap)
	// fmt.Println(files)
	dir := "."
	// ls(dir)
	// a(dir)
	// l(dir)
	// Bfs(dir)
	// r(dir)
    
	// result := functions.LS(dir)
    // fmt.Println(result)
	// test the t flag
	result := functions.T(dir)

	fmt.Println(functions.Filter(result))
}
