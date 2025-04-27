package functions

import (
	"fmt"
	"os"
)

func A(FileName string) []string {
	result := []string{}
	colors := map[string]string{
		"blue":  "\033[94m",
		"reset": "\033[0m",
	}
	fileInfo, err := os.Stat(FileName)
	if err != nil {
		fmt.Println("error in the os.Stat function :", err)
	}
	if fileInfo.IsDir() {
		// fmt.Println(colors["blue"] + "." + colors["reset"])
		// fmt.Println(colors["blue"] + ".." + colors["reset"])
		result = append(result, ".")
		result = append(result, "..")
		content, err1 := os.ReadDir(FileName)
		if err1 != nil {
			fmt.Println("error opening the folder")
			return nil
		}
		for _, x := range content {
			if x.IsDir() {
				temp := colors["blue"] + x.Name() + colors["reset"]
				// fmt.Println(temp)
				result = append(result, temp)
				
			} else {
				result = append(result, x.Name())
				// fmt.Println(x.Name())
			}
		}
	} else {
		// fmt.Println(fileInfo.Name())
		result = append(result, fileInfo.Name())
	}
	return result
}
