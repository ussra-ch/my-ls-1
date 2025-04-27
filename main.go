package main

import (
	"fmt"
	functions "myls1/functions"
	"os"
)

func main() {
	args := os.Args
	Themap, files := functions.ArgumentChecking(args)
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

	if Themap["a"]{
		for _, y := range files{
			result := functions.A(y)
			for _, x := range result{
				fmt.Print(x, " ")
			}
			fmt.Println("")
		}
	}
	// if Themap["ls"]{
	// 	for _, y := range files{
	// 		result := functions.LS(y)
	// 		for _, x := range result{
	// 			fmt.Print(x, "  ")
	// 		}
	// 		fmt.Println("")
	// 	}
	// }
	if Themap["t"]{
		for _, y := range files{
			result := functions.T(y) //mzl kayn hna mouchkil dual . w .., khas htahuma ytratbo bhal l folders/files lakhrin
			for _, x := range result{
				fmt.Print(x.NameTemp, "  ")
			}
			fmt.Println("")
		}
	}

	if Themap["R"]{
		for _, y := range files{
			result := functions.R(y) //makadirch ls 3ad tdkhlul lakhrin. katdkhul nichan l subdir
			for z, x := range result{
				if z == 0 || z == 1{
					continue
				}else{
					for i, t := range x{
						if len(t) == 0{
							continue
						}
						if i == 0{
							fmt.Println(t + ":")
							// fmt.Println("")
						}else{
							fmt.Print(t, "    ")
						}
						// fmt.Print(t, " ")
					}
					fmt.Println("\n")
				}
				}
			// fmt.Println(result)
		}
	}

}

