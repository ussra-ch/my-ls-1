package functions

import (
	"fmt"
	"os"
)

func A(FileName string) {
	colors := map[string]string{
		"blue":  "\033[94m",
		"reset": "\033[0m",
	}
	fileInfo, err := os.Stat(FileName)
	if err != nil {
		fmt.Println("error in the os.Stat function :", err)
	}
	if fileInfo.IsDir() {
		fmt.Println(colors["blue"] + "." + colors["reset"])
		fmt.Println(colors["blue"] + ".." + colors["reset"])
		content, err1 := os.ReadDir(FileName)
		if err1 != nil {
			fmt.Println("error opening the folder")
			return
		}
		for _, x := range content {
			if x.IsDir() {
				temp := colors["blue"] + x.Name() + colors["reset"]
				fmt.Println(temp)
			} else {
				fmt.Println(x.Name())
			}
		}
	} else {
		fmt.Println(fileInfo.Name())
	}
}
