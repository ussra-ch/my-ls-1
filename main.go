package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
)

var  Queue[]string

func main(){
	// args := os.Args
	dir := "."
	// ls(dir)
	// a(dir)
	// l(dir)
	// R(dir)
	Bfs(dir)
}

func ls(path string){
	colors := map[string]string{
		"blue":    "\033[94m",
		"reset":   "\033[0m",
	}
	fileInfo, err := os.Stat(path)
	if err != nil{
		fmt.Println("error in the os.Stat function :", err)
	}
	if fileInfo.IsDir(){
		content, err1 := os.ReadDir(path)
		if err1 != nil{
			fmt.Println("error opening the folder")
			return
		}
		for _, x := range content{
				if x.Name()[0] == '.'{
					continue
				}
				fullPath := filepath.Join(path, x.Name())
				if x.IsDir(){
					Queue = append(Queue, fullPath)
					temp := colors["blue"] + x.Name() + colors["reset"]
					fmt.Print(temp, " ")
				}else{
					fmt.Print(x.Name(), " ")
				}
		}
	}else{
		fmt.Print(fileInfo.Name(), " ")
	}
}

func a(FileName string){
	colors := map[string]string{
		"blue":    "\033[94m",
		"reset":   "\033[0m",
	}
	fileInfo, err := os.Stat(FileName)
	if err != nil{
		fmt.Println("error in the os.Stat function :", err)
	}
	if fileInfo.IsDir(){
		fmt.Println(colors["blue"] + "." + colors["reset"])
		fmt.Println(colors["blue"] + ".." + colors["reset"])
		content, err1 := os.ReadDir(FileName)
		if err1 != nil{
			fmt.Println("error opening the folder")
			return
		}
		for _, x := range content{
				if x.IsDir(){
					temp := colors["blue"] + x.Name() + colors["reset"]
					fmt.Println(temp)
				}else{
					fmt.Println(x.Name())
				}
		}
	}else{
		fmt.Println(fileInfo.Name())
	}
	
}

func l(FileName string){
	FileInfo, err := os.Stat(FileName)
	if err != nil{
		fmt.Println("Error in the 'l' function, Os.stat")
	}
	test, err1 := FileInfo.Sys().(*syscall.Stat_t) //err1 here is a bool
	if !err1{
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

// func pop(queue []string)[]string{
// 	if len(queue) == 0{
// 		return []string{}
// 	}
// 	// res := queue[0]
// 	queue = queue[1:]
// 	return queue
// }

// func R(FileName string) {
// 	fmt.Println(FileName)
// 	ls(FileName)
// 	temp := pop()
// 	if temp == "" {
// 		return
// 	}
// 	fmt.Println(FileName)
// 	R(temp)
// }

func Bfs(DirName string){
	queue := []string{}
	currentPath := ""

	FileInfos, err := os.Stat(DirName)
	if err != nil{
		fmt.Println("Error in Os.stat, Bfs function", err)
	}

	if FileInfos.IsDir(){
		queue = append(queue, DirName)
	}else{
		fmt.Println(DirName)
		return
	}

	for len(queue) > 0{
		currentPath = queue[0]
		queue = queue[1:]
		FileInfos, ok := os.Stat(DirName)
		if ok != nil{
			fmt.Println(ok)
		}
		if !FileInfos.IsDir(){
			fmt.Println(DirName)
			return
		}
		subDir, err := os.ReadDir(currentPath)
		if err != nil{
			fmt.Println("Error in Bfs fucntion, Os.ReadDir", err)
		}
		fmt.Println(currentPath + " :")
		fullPath := ""
		for _, x := range subDir{
			if currentPath[len(currentPath)-1] == '/'{
				fullPath = currentPath + x.Name()
			}else{
				fullPath = currentPath + "/" + x.Name()
			}
			if x.Name()[0] == '.'{
				continue
			}
			fmt.Print(x.Name(), " ")
			
			if x.IsDir(){
				queue = append(queue, fullPath)
			}
		}
		fmt.Println("")
		fmt.Println("")
	}
}


