package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type FlagFields struct {
	Names []string
	Size  int
}

func (f *FlagFields) String() string {
	return fmt.Sprint(f.Names)
}

func (f *FlagFields) Set(v string) error {
	if len(f.Names) > 0 {
		return errors.New("Names len > 0")
	}
	f.Names = append(f.Names, strings.Split(v, ",")...)
	return nil
}

func main() {
	var names FlagFields
	fmt.Println(os.Args)
	if len(os.Args) == 3 {
		names.Names = append(names.Names, os.Args[1])
		names.Size, _ = strconv.Atoi(os.Args[2])
	} else {
		flag.Var(&names, "names", "Comma-separated list")
		flag.Parse()
	}
	fmt.Println(names)
	for _, name := range names.Names {
		if err := ReadText(name, names.Size); err != nil {
			fmt.Println("===========================err======================", err)
			fmt.Println("====================================================")
			fmt.Println("====================================================")
			fmt.Println(err)
			fmt.Println("====================================================")
			fmt.Println("====================================================")
			fmt.Println("===========================err======================")
		}
	}

	GenerateRandom()

}

func ReadText(fileName string, size int) (err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()

	switch size {
	case 0:
		err = ReadString(f)
	default:
		err = ReadSize(f, size)
	}

	return
}

func ReadString(f *os.File) error {
	r := bufio.NewReader(f)
	for {
		text, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("error reading file %s", err)
				return err
			}
		}
		//read line
		fmt.Println(text)
		//read  word
		r := regexp.MustCompile("[^\\s]+")
		words := r.FindAllString(text, -1)
		for i, word := range words {
			fmt.Println(i, word)
		}
	}
	return nil
}

func ReadSize(f *os.File, size int) error {
	buffer := make([]byte, size)
	r := bufio.NewReader(f)
	for {
		n, err := r.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("error reading file %s", err)
				return err
			}
		}
		fmt.Println(string(buffer[:n]))
	}
	return nil
}

func GenerateRandom() {
	file, err := os.Open("/dev/random")
	if err != nil {
		return
	}
	var seed int64
	binary.Read(file, binary.LittleEndian, &seed)
	fmt.Println(seed)
}
