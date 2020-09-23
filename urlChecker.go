package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"mvdan.cc/xurls/v2"
)

func extractURL(str string) []string {
	rxStrict := xurls.Strict()
	foundUrls := rxStrict.FindAllString(str, -1)
	return foundUrls
}

func checkURL(urls []string) {
	for _, v := range urls {
		resp, err := http.Get(v)
		if err != nil {
			print(err.Error())
		} else {
			fmt.Println(string(resp.StatusCode) + resp.Status)
		}
	}
}

func main() {
	content, err := ioutil.ReadFile("urls.txt")
	if err != nil {
		log.Fatal(err)
	}

	textContent := string(content)
	fmt.Println(textContent)

	fmt.Println("--------------------------------------------")

	checkURL(extractURL(textContent))

}
