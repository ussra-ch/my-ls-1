package functions

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

// katprinti hta l files with a dot
func L(FileName string, TheMap map[string]bool, root string) {
		if FileName == "" {
		return
	}
	FileInfo, err := os.Stat(FileName)
	if err != nil {
		return
	}
	test, err1 := FileInfo.Sys().(*syscall.Stat_t) // err1 here is a bool
	if !err1 {
		fmt.Println("syscall struct error, go check it")
	}
	if !TheMap["a"] && FileInfo.Name()[0] == '.' {
		return
	}	
	fmt.Printf("%-10s ", FileInfo.Mode())
	fmt.Printf("%3d ", test.Nlink)
	userName, err2 := user.LookupId(strconv.Itoa(int(test.Uid)))
	if err2 != nil {
		fmt.Println("Error in LookupId, go check it")
	}
	fmt.Printf("%-8s ", userName.Username)
	groupName, err2 := user.LookupGroupId(strconv.Itoa(int(test.Gid)))
	if err2 != nil {
		fmt.Println("Error in LookupId, go check it")
	}
	fmt.Printf("%-8s ", groupName.Name)
	fmt.Printf("%4d ", test.Size)
	time := FileInfo.ModTime()
	fmt.Print(time.Format("Jan 02 15:04"), " ")
	fmt.Println(FileInfo.Name())
}

func CleanPath(fullPath string) string {
	result := ""
	temp := strings.Split(fullPath, "/")
	for i := len(temp) - 1; i >= 0; i-- {
		if len(temp[i]) > 0 && temp[i] != " " {
			result += temp[i]
			return result
		}
	}
	if len(result) == 0 {
		if fullPath != "" {
			result += "/"
		}else{
			result += ".."
		}
	}
	return result
}

func Getwd(path string) (string, int) {
	result := ""
	splitted := strings.Split(path, "/")
	fmt.Println(splitted)
	for i := len(splitted) - 1; i >= 0; i-- { // bach nakhud .
		if splitted[i] != "" {
			return splitted[i], i
		}
	}
	return result, 0
}

func GetRoot(path string, index int) string {
	result := ""
	splitted := strings.Split(path, "/")
	if index == 0{
		return ""
	}
	for i := index-1; i >= 0; i-- { // bach nakhud .
		if splitted[i] != "" {
			return splitted[i]
		}
	}
	return result
}
