package main

import (
	"fmt"
	"os"
	"strings"

	// "strings"

	functions "myls1/functions"
	// "strings"
)

// mzl lina dik total li katkoun m3a ls -l
func main() {
	args := os.Args
	TheMap, files := functions.ArgumentChecking(args)
	// RecursionResult := [][]string{}
	// ReverseResult := []string{}
	// TimeResult := []string{}

	for _, x := range files {
		result := ls(x, TheMap)
		if TheMap["R"] {
			for _, y := range result {
				// full := x + "/" + y

				info, err := os.Stat(y)
				if err != nil {
					fmt.Println("error in the main")
				}
				if y != "" && info.IsDir() {
					Recursion(y, TheMap) // recurse with full path
				}
			}
		}
	}
}

func ls(FileName string, TheMap map[string]bool) []string {
	result := []string{}
	s := strings.Split(FileName, "/")

	if FileName == "." || FileName == "/" {
		if FileName == "." {
		} else {
			FileName = ".."
		}
	} else {
		if s[0] == "" {
			FileName = "." + FileName
		}
		if len(s) == 1 {
			FileName = "./" + FileName
		}
	}

	fileInfos, err := os.Stat(FileName)
	if err != nil {
		fmt.Println("The file name is :", FileName)
		fmt.Println("error in the os.Stat function :", err)
	}

	if fileInfos.IsDir() {
		if TheMap["ls"] {
			result = functions.LS(FileName, TheMap)
		}
		// fmt.Println(functions.Queue)
		if TheMap["t"] {
			// fmt.Println(result)
			us := []string{}
			temp := functions.T(FileName, []string{}, TheMap)
			for _, x := range temp {
				// fmt.Println(x)
				us = append(us, x.NameTemp)
			}
			result = us
			// fmt.Println(result)
			// fmt.Println(result)
		}
		if TheMap["r"] {
			us := []string{}
			// fmt.Println(result)
			us = functions.Reverse("", result, TheMap)
			result = us

			// fmt.Println(result)
		}
		if TheMap["l"] {
			for _, x := range result {
				// fmt.Println(x)
				functions.L(x, TheMap, result[0]) // it always returns the files that start with a dot even if the a flag is false
			}
		}
	} else {
		if TheMap["l"] {
			// fmt.Println(x)
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
		fmt.Println(FileName, " : ")
		rec := ls(FileName, TheMap)
		for r, x := range rec {
			// full := FileName + "/" + x
			Infos, err := os.Stat(x)
			if err != nil {
				fmt.Println("here is the error : ", err, FileName)
				return
			}
			if Infos.IsDir() && Infos.Name() == FileName && r != 0 {
				Recursion(x, TheMap)
			}
		}
	}
}
