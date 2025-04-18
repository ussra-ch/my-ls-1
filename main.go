package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	M "myls1/functions"
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
			provi.name = x.Name()
			provi.timeStamp = FileInfos.ModTime()
			result = append(result, provi)
		}
	} else {
		provi := temp{}
		provi.name = fileInfo.Name()
		provi.timeStamp = fileInfo.ModTime()
		result = append(result, provi)
	}

	new_res := []temp{}

	for _, x := range result {
		new_res = append(new_res, temp{x.name, x.ConvertToUnix()})
	}

	for i := 0; i < len(new_res)-1; i++{
		if new_res[i].timeStamp.(int64) < new_res[i+1].timeStamp.(int64){
			swap := new_res[i]
			new_res[i] = new_res[i+1]
			new_res[i+1] = swap
			i = -1
		}
	}

	for _, x := range new_res {
		fmt.Println(x.String())
	}
}
