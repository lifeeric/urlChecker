package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"mvdan.cc/xurls/v2"
)

func extractURL(str string) []string {
	rxStrict := xurls.Strict()
	foundUrls := rxStrict.FindAllString(str, -1)
	return foundUrls
}

func checkURL(urls []string) {

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, v := range urls {
		go func(v string) {
			defer wg.Done()
			resp, err := http.Head(v)
			if err != nil {
				//print(err.Error())
				fmt.Println("NO RESPONCE!")
			} else {

				switch code := resp.StatusCode; code {
				case 200:
					fmt.Println(v + ": GOOD!")
				case 400, 404:
					fmt.Println(v + ": BAD!")
				default:
					fmt.Println(v + ": UNKNOWN!")

				}

			}
		}(v)

	}

	wg.Wait()
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("help/usage message: To run this program, please pass an argument to it,i.e.: go run urlChecker.go urls.txt")
	} else {

		content, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Fatal(err)

		}

		textContent := string(content)

		fmt.Println("UrlChecker is working now! ")
		fmt.Println("--------------------------------------------------------------------------------")

		checkURL(extractURL(textContent))

	}
}
