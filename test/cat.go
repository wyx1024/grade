package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func printFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		io.WriteString(os.Stdout, scanner.Text()+"\n")

	}
	return nil
}

func main() {
	filename := ""
	arguments := os.Args
	if len(arguments) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}
	for i := 1; i < len(arguments); i++ {
		filename = arguments[i]
		if err := printFile(filename); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
