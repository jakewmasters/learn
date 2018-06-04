package main

import (
	"fmt"
	"os"
)

func main(){
	var path, err = os.Executable()
	if err == nil {
		fmt.Printf("Path to calling executable is %s\n", path)
	}
}