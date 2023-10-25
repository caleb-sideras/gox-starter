package rn

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/caleb-sideras/goxstack/gox/data"
)

const randomOrgAPIURL = "https://www.random.org/integers/?num=1&min=1&max=100&col=1&base=10&format=plain"

func fetchRandomNumber() (int, error) {
	resp, err := http.Get(randomOrgAPIURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var randomNumber int
	_, err = fmt.Sscanf(string(body), "%d", &randomNumber)
	if err != nil {
		return 0, err
	}

	return randomNumber, nil
}

type Number struct {
	Number int
}

func Data(w http.ResponseWriter, r *http.Request) data.PageReturn {
	randomNumber, err := fetchRandomNumber()
	if err != nil {
		log.Println("Error getting random number:", err)
		return data.PageReturn{data.Page{}, err}
	}

	return data.PageReturn{
		Page: data.Page{
			Content:   Number{randomNumber},
			Templates: []string{},
		},
		Error: nil,
	}

}
