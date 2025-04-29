package functions

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)




// done returning slice

type Temp struct {
	NameTemp      string
	TimeStamp any
}

func (t *Temp) ConvertToUnix() int64 {
	return t.TimeStamp.(time.Time).Unix()
}

func T(path string, input[]string)[]Temp {
	// res := []string{}

	result := []Temp{}
	if path != ""{
		fileInfo, err := os.Stat(path)
		if err != nil {
			fmt.Println("error in the os.Stat function :", err)
		}
		if fileInfo.IsDir() {
			content, err1 := os.ReadDir(path)
			if err1 != nil {
				fmt.Println("error opening the folder")
				return nil 
			}
			for _, x := range content {
				//ّ i want to work on all files then ignore the ones that starts with a dot
				fullPath := filepath.Join(path, x.Name()) // it's forbidden to use this function
				provi := Temp{}
				FileInfos, err := os.Stat(fullPath)
				if err != nil {
					fmt.Println("Error in t function, inside the loop. Go check it :) ")
				}
				Queue = append(Queue, fullPath)
				// fmt.Print(temp, " ")
				provi.NameTemp = x.Name()
				provi.TimeStamp = FileInfos.ModTime()
				result = append(result, provi)
			}
		} else {
			// fmt.Print(fileInfo.Name(), " ")
			provi := Temp{}
			provi.NameTemp = fileInfo.Name()
			provi.TimeStamp = fileInfo.ModTime()
			// res = append(res, fileInfo.Name())
			result = append(result, provi)
		}
	}else{
		provi := Temp{}
		// result := []Temp{}
		for _, x := range input{
			FileInfos, err := os.Stat(x)
			if err != nil {
				fmt.Println("error in the os.Stat function :", err)
			}
			provi.NameTemp = x
			provi.TimeStamp = FileInfos.ModTime()
			result = append(result, provi)
		}
	}

	newResult := []Temp{}

	for _, x := range result {
		// fmt.Println(x.timeStamp.(time.Time).Unix()
		newResult = append(newResult, Temp{x.NameTemp, x.ConvertToUnix()})
		// fmt.Println(newResult)
	}

	for i := 0; i < len(newResult)-1; i++ {
		if newResult[i].TimeStamp.(int64) < newResult[i+1].TimeStamp.(int64) {
			swap := newResult[i]
			newResult[i] = newResult[i+1]
			newResult[i+1] = swap
			i = -1
		}
	}
	return newResult
}

func (t *Temp) String() string {
	return fmt.Sprintf("name: %v and time %v", t.NameTemp, t.TimeStamp)
}




// utility to only show the string as the final result //
// mashi muhima just for debugging goals
func  Filter(temps []Temp) string {
	result := ""
    for i:=0; i< len(temps); i++ {
    result+= fmt.Sprintf("%v ", temps[i].NameTemp)
	} 
    return result
}
