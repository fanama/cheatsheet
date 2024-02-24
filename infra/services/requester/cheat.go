package requester

import (
	"fmt"
	"io"
	"net/http"

	"jaytaylor.com/html2text"
)

type Cheat struct {
	url string
}

func BuildCheat() Cheat {
	return Cheat{url: "http://cheat.sh/"}
}

func (this *Cheat) Ask(field string, input string) (string, error) {

	resp, err := http.Get(this.url + field + "/" + input)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	text, err := html2text.FromString(string(body), html2text.Options{PrettyTables: true})

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return text, nil
}
