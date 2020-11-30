package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetWordData(word string) ([]Entry, error) {
	path := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)
	resp, err := http.Get(path)
	if err != nil {
		fmt.Println("err didn't hit api")
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("didn't like what you got back")
		return nil, err
	}

	keysBody := []byte(`[{"id": 1,"key": "-"},{"id": 2,"key": "-"},{"id": 3,"key": "-"}]`)

	dic := dictionaryResponse{}


	err = json.Unmarshal(body, &dic)
	if err != nil {
		fmt.Println("didn't unmarhsal")
		return nil, err
	}

	return dic.Response, nil

}

type Entry struct {
	Word      string      `json:"word,omitEmpty"`
	Phonetics []Phonetics `json:"phonetics,omitEmpty"`
	Origin    string      `json:"origin,omitEmpty"`
	Meanings  Meaning     `json:"meanings,omitEmpty"`
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
	Response []Entry
}
