package functions

import (
	"fmt"
	"os"
	"path/filepath"
)



// done returning slice
func LS(path string, TheMap map[string]bool)  []string{
	result := []string{}
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("error in the os.Stat function :", err)
	}
	if fileInfo.IsDir() {
		content, err1 := os.ReadDir(path)
		if err1 != nil {
			fmt.Println("error opening the folder") 
			return nil
		}
		if TheMap["a"]{
			result = append(result, ".", "..")
		}
		for _, x := range content {
			if x.Name()[0] == '.' && !TheMap["a"]{
				continue
			}
			// this to be changed after ///////////////////////////////////////////////
			fullPath := filepath.Join(path, x.Name())
			if x.IsDir() {
				Queue = append(Queue, fullPath)
			}
			result = append(result, fullPath)
		}
	} else {
		result = append(result, fileInfo.Name())
	}

	return result
}
