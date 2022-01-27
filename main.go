package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"jaytaylor.com/html2text"
)

func main() {

	fmt.Println("===========================================================")
	if len(os.Args) < 2 {
		fmt.Println("not enough argument")
		os.Exit(1)
	}

	field := os.Args[1]
	topic := ""

	if len(os.Args) > 2 {
		topic = strings.Join(os.Args[2:], "+")
	}

	fmt.Println("===========================================================")
	fmt.Println("Getting ", field, "/", topic+"...")

	resp, err := http.Get("http://cheat.sh/" + field + "/" + topic)

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	text, err := html2text.FromString(string(body), html2text.Options{PrettyTables: true})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(text)

}
