package functions

import (
	"fmt"
	"os"
	"path/filepath"
)

// Old implementation of the r flag
func Reverse(path string, input []string) []string{
	colors := map[string]string{
		"blue":  "\033[94m",
		"reset": "\033[0m",
	}
	res := []string{}
	final := []string{}
	if path != ""{
		fileInfo, err := os.Stat(path)
		if err != nil {
			fmt.Println("error in the os.Stat function :", err)
			return []string{}
		}
		if fileInfo.IsDir() {
			content, err1 := os.ReadDir(path)
			if err1 != nil {
				fmt.Println("error opening the folder")
				return []string{}
			}
			for _, x := range content {
				if x.Name()[0] == '.' {
					continue
				}
				// to be changed after 
				fullPath := filepath.Join(path, x.Name())
				if x.IsDir() {
					Queue = append(Queue, fullPath)
					temp := colors["blue"] + x.Name() + colors["reset"]
					// fmt.Print(temp, " ")
					res = append(res, temp)
				} else {
					// fmt.Print(x.Name(), " ")
					res = append(res, x.Name())
				}
			}
		} else {
			// fmt.Print(fileInfo.Name(), " ")
			res = append(res, fileInfo.Name())
		}
	}else{
		res = input
	}
	

	for i := len(res) - 1; i >= 0; i-- {
		// fmt.Print(res[i], " ")
		final = append(final, res[i])
	}
	return final
}
