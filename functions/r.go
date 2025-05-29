package functions

import (
	"fmt"
	"os"
	"path/filepath"
)

// Old implementation of the r flag
func Reverse(path string, input []string, TheMap map[string]bool) []string{
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
				if x.Name()[0] == '.' && !TheMap["a"]{
					continue
				}
				// to be changed after //////////////////////////
				fullPath := filepath.Join(path, x.Name())
				if x.IsDir() {
					Queue = append(Queue, fullPath)
					res = append(res, fullPath)
				} else {
					res = append(res, fullPath)
				}
			}
		} else {
			res = append(res, path)
		}
	}else{
		res = input
	}
	

	for i := len(res) - 1; i >= 0; i-- {
		final = append(final, res[i])
	}
	return final
}
