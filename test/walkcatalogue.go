package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var F = false
var D = false

func main() {
	path := ""
	arg1 := flag.Bool("d", false, "ddd")
	arg2 := flag.Bool("f", false, "fff")
	flag.Parse()
	if len(flag.Args()) == 1 {
		args := flag.Args()
		path = args[0]
	}

	F = *arg2
	D = *arg1
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		fileinfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		if fileinfo.IsDir() && D {
			fmt.Println("+", path)
		}

		if fileinfo.Mode().IsRegular() && F {
			fmt.Println("*", path)
		}
		return nil
	})

}
