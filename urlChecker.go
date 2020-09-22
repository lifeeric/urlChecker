package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("urls.txt")
	if err != nil {
		log.Fatal(err)
	}

	textContent := string(content)
	fmt.Println(textContent)

}
