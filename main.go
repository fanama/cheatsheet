package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fanama/cheatsheet/infra/services/requester"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("===========================================================")
		fmt.Println("not enough argument")
		os.Exit(1)
	}

	field := os.Args[1]
	topic := ""

	if len(os.Args) > 2 {
		topic = strings.Join(os.Args[2:], "+")
	}

	cheatController := requester.BuildCheat()

	text, err := cheatController.Ask(field, topic)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n \n", text)

}
