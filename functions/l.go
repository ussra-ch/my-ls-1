package functions

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	// "path/filepath"
)

// katprinti hta l files with a dot
func L(FileName string, TheMap map[string]bool, root string) {
	// fmt.Println(FileName)
	s:=strings.Split(FileName,"/")
	// fmt.Println(len(s))
	temp := []string{}
	for i := 0; i < len(s); i++{
		if s[i] != ""{
			temp = append(temp, s[i])
		}
	}
	// fmt.Println(len(temp))
	if len(s) == 1||s[0]==""{
		FileName="./"+FileName
	}
	// splitted := strings.Split(FileName, "/")
	// fmt.Println(splitted)
	// if len(splitted) == 1{
	// 	splitted[0]="."
	// 	splitted=append(splitted,root )
		
	// }
	FileInfo, err := os.Stat(FileName)
	// fmt.Println("File infos:", FileInfo.Name())
	if err != nil {
		fmt.Println(FileName)
		fmt.Println("Error in the 'l' function, Os.stat", err)
	}
	test, err1 := FileInfo.Sys().(*syscall.Stat_t) // err1 here is a bool
	if !err1 {
		fmt.Println("syscall struct error, go check it")
	}
	if !TheMap["a"] && FileInfo.Name()[0] == '.' {
		return
	}
	// fmt.Println(FileName)

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
	fmt.Printf("%s ", groupName.Name)
	fmt.Printf("%d ", test.Size)
	time := FileInfo.ModTime()
	fmt.Print(time.Format("Jan 02 15:04"), " ")
	//fmt.Println(FileInfo.Name()) // January 2nd, 2006 at 3:04:05 PM (MST) — is Go’s reference time.
	// index := len(s)-1
	// state := 0
	// for i := index; i >= 0; i--{
	// 	if temp[i] != ""{
	// 		if index != 0{
	// 			state = 2
	// 		}else{
	// 			state = 1
	// 		}
	// 		index = i
	// 		FileInfo, err = os.Stat(temp[i])
	// 		// fmt.Println("File infos:", FileInfo.Name())
	// 		if err != nil {
	// 			fmt.Println(temp[i])
	// 			fmt.Println("Error in the 'l' function, Os.stat", err)
	// 		}
	// 		break
	// 	}
	// }

	// if FileInfo.IsDir()&& state == 1{
	// 	fmt.Println("..")
	// }else if FileInfo.IsDir()&& state == 2{
	// 	fmt.Println(".")
	// }else{
	// 	fmt.Println(FileInfo.Name())
	// }

	// fmt.Println("Leeeenght is :", len(s))
	fmt.Println("file name is :", FileName)
	// fmt.Println("slice content is :", s)
	if s[len(s)-1] == FileName{
		fmt.Println("..")
	}else if s[0] == "."{
		fmt.Println(".")
	}else{
		fmt.Println(FileInfo.Name())
	}
	// if FileInfo.Name()==temp[1]{
	// 	fmt.Println(".")
	// }else{
	// 	fmt.Println(CleanPath(FileInfo.Name()))
	// }
	// if FileInfo.Name() == splitted[1] {
		
	// 	fmt.Println(".")
	// } else if  FileInfo.Name() == splitted[0]{
	// 	fmt.Println("..")
	// }else{
	// 	fmt.Println(CleanPath(FileInfo.Name()))
	// }
	
	// 	abs, _ := filepath.Abs(FileName)
	// fmt.Println("here",abs)
}

func CleanPath(fullPath string) string {
	result := ""
	temp := strings.Split(fullPath, "/")

	for i := len(temp) - 1; i >= 0; i-- {
		// fmt.Println(temp[i])
		if len(temp[i]) > 0 && temp[i] != " " {
			// fmt.Println("111")
			result += temp[i]
			return result
		}
	}
	if len(result) == 0 {
		result += "/"
	}

	return result
}

func Getwd(path string) string {
	result := ""
	splitted := strings.Split(path, "/")
	fmt.Println(splitted)
	for i := len(splitted) - 1; i >= 0; i-- { // bach nakhud .
		fmt.Println("111")
		if splitted[i] != "" {
			fmt.Println("222")
			return splitted[i]
		}
	}
	return result
}

func GetRoot(path string) string {
	result := ""
	splitted := strings.Split(path, "/")
	// fmt.Println(splitted)
	for i := len(splitted) - 2; i >= 0; i-- { // bach nakhud .
		if splitted[i] != "" {
			return splitted[i]
		}
	}
	return result
}
