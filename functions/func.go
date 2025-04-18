package functions

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

var Queue []string

func ls(path string) {
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
			return
		}
		for _, x := range content {
			if x.Name()[0] == '.' {
				continue
			}
			fullPath := filepath.Join(path, x.Name())
			if x.IsDir() {
				Queue = append(Queue, fullPath)
				temp := colors["blue"] + x.Name() + colors["reset"]
				fmt.Print(temp, " ")
			} else {
				fmt.Print(x.Name(), " ")
			}
		}
	} else {
		fmt.Print(fileInfo.Name(), " ")
	}
}

func a(FileName string) {
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

func l(FileName string) {
	FileInfo, err := os.Stat(FileName)
	if err != nil {
		fmt.Println("Error in the 'l' function, Os.stat")
	}
	test, err1 := FileInfo.Sys().(*syscall.Stat_t) // err1 here is a bool
	if !err1 {
		fmt.Println("syscall struct error, go check it")
	}

	fmt.Print(FileInfo.Mode(), "  ")
	fmt.Print(test.Nlink, " ")
	userName, err2 := user.LookupId(strconv.Itoa(int(test.Uid)))
	if err2 != nil {
		fmt.Println("Error in LookupId, go check it")
	}
	fmt.Print(userName.Username, " ")
	groupName, err2 := user.LookupGroupId(strconv.Itoa(int(test.Gid)))
	if err2 != nil {
		fmt.Println("Error in LookupId, go check it")
	}
	fmt.Print(groupName.Name, " ")
	fmt.Print(test.Size, " ")
	time := FileInfo.ModTime()
	fmt.Print(time.Format("Jan 02 15:04"), " ") // January 2nd, 2006 at 3:04:05 PM (MST) — is Go’s reference time.
	fmt.Println(FileInfo.Name())
}

func Bfs(DirName string) {
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

func argument_check(args []string) (map[string]bool, []string) {
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
	}
	if len(files) == 0 {
		files = append(files, ".")
	}
	return TheMap, files
}

func r(path string) {
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
