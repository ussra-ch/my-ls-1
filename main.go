package main

import (
	"fmt"
	"os"
	// "strings"

	functions "myls1/functions"
)

// var isrot bool

// mzl lina dik total li katkoun m3a ls -l
func main() {
	args := os.Args
	TheMap, files := functions.ArgumentChecking(args)

	for _, x := range files {
		result := ls(x, TheMap)
		if TheMap["R"] {
			for _, y := range result {
				info, err := os.Stat(y)
				if err != nil {
				}
				if info.IsDir() {
					Recursion(y, TheMap) // recurse with full path
				}
			}
		}
	}
}

func ls(FileName string, TheMap map[string]bool) []string {
	result := []string{}
	if TheMap["R"]{
		fmt.Print(FileName)
		fmt.Println(":")
	}
	// s := strings.Split(FileName, "/")

	// if FileName == "." || FileName == "/" {
	// 	if FileName != "." {
	// 		FileName = ".."
	// 	}
	// } else {
	// 	if s[0] == "" {
	// 		// FileName = "." + FileName
	// 	}
	// 	if len(s) == 1 {
	// 		// FileName = "./" + FileName
	// 	}
	// }

	fileInfos, err := os.Stat(FileName)
	if err != nil {
		fmt.Printf("ls: cannot access '%s': No such file or directory", FileName)
		return nil
	}

	if fileInfos.IsDir() {
		if TheMap["ls"] {
			result = functions.LS(FileName, TheMap)
		}
		if TheMap["t"] {
			us := []string{}
			temp := functions.T(FileName, []string{}, TheMap)
			for _, x := range temp {
				us = append(us, x.NameTemp)
			}
			result = us
		}
		if TheMap["r"] {
			us := []string{}
			us = functions.Reverse("", result, TheMap)
			result = us
		}
		if TheMap["l"] {
			total := int64(0)
			block := int64(512)
			for _, x := range result{
				fileInfos, err := os.Stat(x)
				if err != nil{
					continue
				}
				if fileInfos.IsDir(){
					continue
				}
				temp := (fileInfos.Size() + block-1)/block
				total += temp
			}
			fmt.Println("total", total)
			for _, x := range result {
				functions.L(x, TheMap, result[0])
				// if !TheMap["a"] {

				// 	functions.L(x, TheMap, result[0])
				// } else if TheMap["a"] {
				// 	functions.L(x, TheMap, result[0])
				// }
			}
		}
	} else {
		if TheMap["l"] {
			functions.L(FileName, TheMap, "") // it always returns the files that start with a dot even if the a flag is false
		} else {
			fmt.Print(fileInfos.Name())
		}
	}
	if !TheMap["l"] {
		for _, x := range result {
			if len(x) != 0 && !TheMap["a"] && x[0] == '.' {
				continue
			} else if len(x) != 0 {
				fmt.Print(functions.CleanPath(x), "    ")
			}
		}
		fmt.Println("")
	}
	return result
}

// it still does not work with .. directory
func Recursion(FileName string, TheMap map[string]bool) {
	if FileName == "." || FileName == ".." {
		return
	}
	FileInfos, err := os.Stat(FileName)
	if err != nil {
		fmt.Println("here here")
		return
	}
	if !FileInfos.IsDir() {
		return
	} else {
		fmt.Println("")
		// fmt.Println(FileName, " : ")
		rec := ls(FileName, TheMap)
		for _, x := range rec {
			if x == "" {
				continue
			}
			Infos, err := os.Stat(x)
			if err != nil {
				fmt.Println("here is the error : ", err, x)
				return
			}

			if Infos.IsDir(){
				Recursion(x, TheMap)
			}

		}
	}
}
