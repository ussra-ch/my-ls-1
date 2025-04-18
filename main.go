package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	M "myls1/math"
)

func main() {
	// args := os.Args
	// Themap, files := argument_check(args)
	// fmt.Println(args)
	// fmt.Println(Themap)
	// fmt.Println(files)
	dir := "."
	// ls(dir)
	// a(dir)
	// l(dir)
	// Bfs(dir)
	// r(dir)
	t(dir)
}

type temp struct {
	name      string
	timeStamp any
}

func (t *temp) String() string {
	return fmt.Sprintf("name: %v and time %v", t.name, t.timeStamp)
}

func (t *temp) ConvertToUnix() int64 {
	return t.timeStamp.(time.Time).Unix()
}

func t(path string) {
	// res := []string{}

	result := []temp{}
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
			//ّ i want to work on all files then ignore the ones that starts with a dot
			fullPath := filepath.Join(path, x.Name())
			provi := temp{}
			FileInfos, err := os.Stat(fullPath)
			if err != nil {
				fmt.Println("Error in t function, inside the loop. Go check it :) ")
			}
			M.Queue = append(M.Queue, fullPath)
			// fmt.Print(temp, " ")
			provi.name = x.Name()
			provi.timeStamp = FileInfos.ModTime()
			result = append(result, provi)
		}
	} else {
		// fmt.Print(fileInfo.Name(), " ")
		provi := temp{}
		provi.name = fileInfo.Name()
		provi.timeStamp = fileInfo.ModTime()
		// res = append(res, fileInfo.Name())
		result = append(result, provi)
	}

	new_res := []temp{}

	for _, x := range result {
		// fmt.Println(x.timeStamp.(time.Time).Unix()
		new_res = append(new_res, temp{x.name, x.ConvertToUnix()})
		fmt.Println(new_res)

	}

	fmt.Println("new_res", new_res)
	// fmt.Println(time.Now())
	// fmt.Println(result)
	// for i := 0; i < len(result)-1; i++{
	// 	if result[i].timeStamp < result[i+1].timeStamp{
	// 		swap := result[i]
	// 		result[i] = result[i+1]
	// 		result[i+1] = swap
	// 		i = -1
	// 	}
	// }

	for _, x := range new_res {
		fmt.Println(x.String())
	}
}
