package main

import (
	"fmt"
	functions "myls1/functions"
	"os"
	// "strings"
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
				// full := x + "/" + y 
				// fmt.Println(y)
				info, err := os.Stat(y)
				if err != nil {
					fmt.Println("error in the main")
				}
				if info.IsDir() {
					Recursion(y, TheMap)         // recurse with full path
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
			result = functions.LS(FileName, TheMap)
		}
		if TheMap["t"]{
			// fmt.Println(result)
			us := []string{}
			temp := functions.T(FileName, []string{}, TheMap)
			for _, x := range temp{
				// fmt.Println(x)
				us = append(us, x.NameTemp)
			}
			result = us
			// fmt.Println(result)
			// fmt.Println(result)
		}
		if TheMap["r"]{
			us := []string{}
			// fmt.Println(result)
			us = functions.Reverse("", result, TheMap)
			result = us
			
			// fmt.Println(result)
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
				fmt.Print(functions.CleanPath(x), "    ")
			}
		}
		fmt.Println("")
	}
	return result
}

//it still does not work with .. directory
func Recursion(FileName string, TheMap map[string]bool){
	if FileName == "." || FileName == ".."{
		return
	}
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
			// full := FileName + "/" + x
			Infos, err := os.Stat(x)
			if err != nil{
				fmt.Println("here is the error : ", err, FileName)
				return
			}
			if Infos.IsDir(){
				Recursion(x, TheMap)
			}
		}
	}
}

