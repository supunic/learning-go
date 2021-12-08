package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(content))

	if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
		log.Fatalln(err)
	}

	r := bytes.NewBuffer([]byte("abc"))
	content2, _ := ioutil.ReadAll(r)
	fmt.Println(string(content2))
}
