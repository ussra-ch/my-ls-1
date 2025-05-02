package functions

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
	"strings"
)

//katprinti hta l files with a dot
func L(FileName string)  {
	FileInfo, err := os.Stat(FileName)
	if err != nil {
		fmt.Println("Error in the 'l' function, Os.stat", err)
	}
	test, err1 := FileInfo.Sys().(*syscall.Stat_t) // err1 here is a bool
	if !err1 {
		fmt.Println("syscall struct error, go check it")
	}

	// if !TheMap["a"] && FileInfo.Name()[0] != '.'{
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
		fmt.Println(CleanPath(FileInfo.Name()))
	// }
}



func CleanPath(fullPath string)string{
	result := ""
	temp := strings.Split(fullPath, "/")

	for i := len(temp)-1; i >= 0; i--{
		// fmt.Println(temp[i])
		if len(temp[i]) > 0 && temp[i] != " "{
			// fmt.Println("111")
			result += temp[i]
			return result
		}
	}
	if len(result) == 0{
		result += "/"
	}

	return result
}