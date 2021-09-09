package main

import (
	"bufio"
	"fmt"
	"go-growth/cache"
	"os"
	"strings"
)

func main() {
	err := cache.LOAD()
	if err != nil {
		fmt.Println(err.Error())
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)
		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "PRINT":
			cache.PRINT()
		case "STOP":
			err := cache.Save()
			if err != nil {
				fmt.Println("STOP", err.Error())
			}
			return
		case "DELETE":
			if !cache.DELETE(tokens[1]) {
				fmt.Println("Delete operations failed")
			}
		case "ADD":
			n := cache.MyElement{tokens[2], tokens[3], tokens[4]}
			if !cache.ADD(tokens[1], n) {
				fmt.Println("Add operation failed")
			}
		case "LOOKUP":
			n := cache.LOOKUP(tokens[1])
			if n != nil {
				fmt.Printf("%v\n", n)
			}

		case "CHANGE":
			n := cache.MyElement{tokens[2], tokens[3], tokens[4]}
			if !cache.CHANGE(tokens[1], n) {
				fmt.Println("Update operation failed")
			}

		default:
			fmt.Println("Unknown command - please try again!")
		}
	}
}
