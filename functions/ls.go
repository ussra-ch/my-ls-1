package functions

import (
	"fmt"
	"os"
	"path/filepath"
)




// done returning slice
func LS(path string)  []string{
	result := []string{}
	colors := map[string]string{
		"blue":  "\033[94m",
		"reset": "\033[0m",
	}
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
		for _, x := range content {
			if x.Name()[0] == '.' {
				continue
			}
			// this to be changed after 
			fullPath := filepath.Join(path, x.Name())
			if x.IsDir() {
				Queue = append(Queue, fullPath)
				temp := colors["blue"] + x.Name() + colors["reset"]
				// fmt.Print(temp, "  ")
				result = append(result,  temp)
			} else {	
				// _, err:= fmt.Print(x.Name(), "  ")
				// if err!= nil {
				// 	return nil
				// }
				result = append(result, x.Name())
			}
		}
	} else {
		// fmt.Print(fileInfo.Name(), "  ")
		result = append(result, fileInfo.Name())
	}

	return result
}
