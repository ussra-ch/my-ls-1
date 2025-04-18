package functions

import (
	"fmt"
	"os"
	"path/filepath"
)

// BFS algorithm
func R(DirName string) {
	queue := []string{}
	currentPath := ""

	FileInfos, err := os.Stat(DirName)
	if err != nil {
		fmt.Println("Error in Os.stat, Bfs function", err)
	}

	if FileInfos.IsDir() {
		queue = append(queue, DirName)
	} else {
		fmt.Println(DirName)
		return
	}

	for len(queue) > 0 {
		currentPath = queue[0]
		queue = queue[1:]
		FileInfos, ok := os.Stat(DirName)
		if ok != nil {
			fmt.Println(ok)
		}
		if !FileInfos.IsDir() {
			fmt.Println(DirName)
			return
		}
		subDir, err := os.ReadDir(currentPath)
		if err != nil {
			fmt.Println("Error in Bfs fucntion, Os.ReadDir", err)
		}
		fmt.Println(currentPath + " :")
		fullPath := ""
		for _, x := range subDir {
			if currentPath[len(currentPath)-1] == '/' {
				fullPath = currentPath + x.Name()
			} else {
				fullPath = currentPath + "/" + x.Name()
			}
			if x.Name()[0] == '.' {
				continue
			}
			fmt.Print(x.Name(), " ")

			if x.IsDir() {
				queue = append(queue, fullPath)
			}
		}
		fmt.Println("")
		fmt.Println("")
	}
}

// Old implementation of the r flag

func Reverse(path string) {
	colors := map[string]string{
		"blue":  "\033[94m",
		"reset": "\033[0m",
	}
	res := []string{}
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("error in the os.Stat function :", err)
	}
	if fileInfo.IsDir() {
		content, err1 := os.ReadDir(path)
		if err1 != nil {
			fmt.Println("error opening the folder")
			return
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

	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i], " ")
	}

}
