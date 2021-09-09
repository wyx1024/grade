package cache

import (
	"encoding/json"
	"fmt"
	"os"
)

type MyElement struct {
	Name    string `json:"name"`
	SurName string `json:"sur_name"`
	ID      string `json:"id"`
}

var DATA = make(map[string]MyElement)
var DATAFILE = "./../cache/data.json"

func LOOKUP(k string) *MyElement {
	if _, ok := DATA[k]; ok {
		n := DATA[k]
		return &n
	}
	return nil
}

func ADD(k string, v MyElement) bool {
	if k == "" {
		return false
	}

	if LOOKUP(k) != nil {
		return false
	}

	DATA[k] = v

	return true
}

func DELETE(k string) bool {
	if LOOKUP(k) == nil {
		return false
	}
	delete(DATA, k)
	return true
}

func CHANGE(k string, v MyElement) bool {
	DATA[k] = v
	return true
}

func PRINT() {
	for k, element := range DATA {
		fmt.Printf("Key:%s, Value:%v\n", k, element)
	}
}

func Save() error {
	fmt.Println("Saveing ", DATAFILE)
	if err := os.Remove(DATAFILE); err != nil {
		fmt.Println(err.Error())
	}
	savaTo, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println("Cannot create ", DATAFILE)
		return err
	}
	defer savaTo.Close()
	encoder := json.NewEncoder(savaTo)
	if err = encoder.Encode(DATA); err != nil {
		return err
	}
	return nil
}

func LOAD() error {
	fmt.Println("Loading ", DATAFILE)
	loadForm, err := os.Open(DATAFILE)
	if err != nil {
		fmt.Println("Empty key/value store!")
		return err
	}
	decoder := json.NewDecoder(loadForm)
	decoder.Decode(&DATA)
	return nil

}
