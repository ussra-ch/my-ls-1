package main

import (
	"fmt"
	functions "myls1/functions"
	"os"
	"strings"
)
// mzl lina dik total li katkoun m3a ls -l
func main() {
	args := os.Args
	TheMap, files := functions.ArgumentChecking(args)
	// RecursionResult := [][]string{}
	// ReverseResult := []string{}
	// TimeResult := []string{}

	for _, x := range files{
		result := ls(x, TheMap)
		if TheMap["R"]{
			for _, y := range result{
				full := x + "/" + y 
				info, err := os.Stat(full)
				if err != nil {
					fmt.Println("error in the main")
				}
				if info.IsDir() {
					Recursion(full, TheMap)         // recurse with full path
				}
			}
		}
	}
}


func ls(FileName string, TheMap map[string]bool) []string{
	result := []string{}
	fileInfos, err := os.Stat(FileName)
	if err != nil {
		fmt.Println("error in the os.Stat function :", err)
	}

	if fileInfos.IsDir(){
		if TheMap["ls"]{
			result = functions.LS(FileName)
		}
		if TheMap["t"]{
			us := []string{}
			temp := functions.T(FileName, []string{})
			for _, x := range temp{
				us = append(us, CleanPath(x.NameTemp))
			}
			result = us
		}
		if TheMap["r"]{
			us := []string{}
			us = functions.Reverse("", result)
			result = us
		}
		if TheMap["l"]{
			for _, x := range result{
				functions.L(x) //it always returns the files that start with a dot even if the a flag is false
			}
		}
	}
	if !TheMap["l"]{
		for _, x := range result{
			if !TheMap["a"] && x[0] == '.'{
				continue
			}else{
				fmt.Print(x, "    ")
			}
		}
		fmt.Println("")
	}
	return result
}

//it still does not work with .. directory
func Recursion(FileName string, TheMap map[string]bool){
	FileInfos, err := os.Stat(FileName)
	if err != nil{
		fmt.Println("here here")
		return
	}
	if !FileInfos.IsDir(){
		return
	}else{
		fmt.Println("")
		fmt.Println(FileName, " : ")
		rec := ls(FileName, TheMap)
		for _, x := range rec{
			full := FileName + "/" + x
			Infos, err := os.Stat(full)
			if err != nil{
				fmt.Println("here is the error")
				return
			}
			if Infos.IsDir(){
				Recursion(full, TheMap)
			}
		}
	}
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