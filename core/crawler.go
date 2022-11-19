package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func freeDictionaryApi(word string) (Word, error) {
	// send a request to the dictionary website
	url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if resp.Status != "200 OK" {
		return Word{}, fmt.Errorf("Word not found")
	}

	defer resp.Body.Close()
	// parse the response JSON to Word
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var words []Word
	err = json.Unmarshal(body, &words)
	if err != nil {
		panic(err)
	}
	return words[0], nil
}

func Crawl(spell string) (Word, error) {
	w, err := freeDictionaryApi(spell)
	if err != nil {
		return Word{}, err
	}
	// Add more dictionary crawlers in the future
	return w, nil
}
