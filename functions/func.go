package functions

import (
	"fmt"
	"strings"
)





var Queue []string

func ArgumentChecking(args []string) (map[string]bool, []string) {
	TheMap := map[string]bool{}
	files := []string{}

	TheMap["ls"] = false
	TheMap["R"] = false
	TheMap["r"] = false
	TheMap["l"] = false
	TheMap["a"] = false
	TheMap["t"] = false
	TheMap["all"] = false
	TheMap["reverse"] = false
	TheMap["time"] = false
	TheMap["recursive"] = false

	if args[1] != "ls" {
		fmt.Println("Please enter the ls command")
		return map[string]bool{}, []string{}
	} else {
		TheMap["ls"] = true
	}

	if len(args) == 2 {
		TheMap["ls"] = true
		files = append(files, ".")
		return TheMap, files
	}

	for i, x := range args {
		if x[0] == '-' && x[1] == '-' {
			TheMap[string(x[2:])] = true
		} else if x[0] == '-' {
			if len(x) > 2 {
				temp := strings.Split(x, "")
				for _, y := range temp {
					TheMap[y] = true
				}
			} else {
				TheMap[string(x[1])] = true
			}
		} else if x != "ls" && i != 0 {
			files = append(files, x)
		}
		if x == "--recursive"{
			TheMap["R"] = true
		}
		if x == "--time"{
			TheMap["t"] = true
		}
		if x == "--reverse"{
			TheMap["r"] = true
		}
		if x == "--all"{
			TheMap["l"] = true
			TheMap["R"] = true
			TheMap["r"] = true
			TheMap["a"] = true
			TheMap["t"] = true
		}
		// fmt.Println(x)
	}
	if len(files) == 0 {
		files = append(files, ".")
	}
	
	return TheMap, files
}
