package main

import{
	"fmt"
	"os"
	"path/filepath"
}

func main(){
	filepath.Walk(".", func(path string, infoos.FileInfo, err error) error{
		fmt.Println(path)
		return nil
	})
}