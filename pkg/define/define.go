package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetWordData(lang, word string) ([]Entry, error) {
	path := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/%s/%s", lang, word)
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	dic := dictionaryResponse{}

	err = json.Unmarshal(body, &dic)
	if err != nil {
		return nil, err
	}

	return dic.Response, nil

}

type Entry struct {
	Word      string    `json:"word,omitEmpty"`
	Phonetics Phonetics `json:"phonetics,omitEmpty"`
	Origin    string    `json:"origin,omitEmpty"`
	Meanings  Meaning   `json:"meanings,omitEmpty"`
}

type Phonetics struct {
	Text  string `json:"text,omitEmpty"`
	Audio string `json:"audio,omitEmpty"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech,omitEmpty"`
	Definitions  []Definition `json:"definitions,omitEmpty"`
}

type Definition struct {
	Def      string   `json:"definition,omitEmpty"`
	Example  string   `json:"example,omitEmpty"`
	Synonyms []string `json:"synonyms,omitEmpty"`
}

type dictionaryResponse struct {
	Response []Entry `json:"omitEmpty"`
}
