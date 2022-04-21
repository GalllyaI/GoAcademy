package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Bartender struct {
	url    string
	Drinks []Drink
}

type Drink struct {
	Name            string `json:"drinks,omitempty"`
	StrInstructions string `json:"strInstructions,omitempty"`
}

func (b *Bartender) Start() string {
	fmt.Println("What would you like to drink")
	var drink string
	fmt.Scanln(&drink)

	return drink
}

func GetRecipe(b *Bartender) {

	drink := ""
	for {
		drink = b.Start()
		if drink == "nothing" {
			fmt.Println("Have a nice day!")
			break
		}

		u, err := url.Parse(b.url)

		if err != nil {
			fmt.Println(err)
		}

		q := u.Query()
		q.Add("s", drink)
		u.RawQuery = q.Encode()
		req, err := http.NewRequest("GET", u.String(), nil)

		if err != nil {
			fmt.Println(err)
		}

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			fmt.Println(err)
		}

		defer res.Body.Close()
		resp := b
		json.NewDecoder(res.Body).Decode(&resp)

		if len(resp.Drinks) == 0 {
			fmt.Println("not available now")
			fmt.Println()
			continue
		}

		// output := strings.Join(strings.Split(resp.Drinks[0].StrInstructions, "."), ".\n")

		output := strings.Split(resp.Drinks[0].StrInstructions, ".")
		for ind, str := range output {
			fmt.Printf("%d. %s.\n", ind+1, str)
		}

		fmt.Println()
	}

}

func main() {
	url := "https://www.thecocktaildb.com/api/json/v1/1/search.php"
	b := &Bartender{url: url}
	GetRecipe(b)
}
