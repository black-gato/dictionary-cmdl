package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetWordData(lang, word string) ([]*Entry, error) {
	path := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/%s/%s", lang, word)
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

	dic := &dictionaryResponse{}

	err = json.Unmarshal(body, &dic.Response)
	if err != nil {
		fmt.Println("didn't unmarhsal")
		return nil, err
	}

	return dic.Response, nil

}

func GetEntry(lang, word string) (string, error) {
	data, err := GetWordData(lang, word)
	if err != nil {
		return "nil", err
	}

	dic := make([]CommandDefinition, 20)
	var start, total int

	for ml := range data[0].Meanings {

		total += len(data[0].Meanings[ml].Definitions)

		for _, def := range data[0].Meanings[ml].Definitions {
			for i := start; i < total; i++ {

				dic[i].Word = word
				dic[i].POS = data[0].Meanings[ml].PartOfSpeech
				dic[i].Def = def.Def
			}

		}

		start += total

	}

	entry := fmt.Sprintf("Word:			%s\n\nPart of Speech:		%s\n\nDef:			%s\n\n", dic[0].Word, dic[0].POS, dic[0].Def)

	return entry, err

}

type Entry struct {
	Word      string      `json:"word,omitEmpty"`
	Phonetics []Phonetics `json:"phonetics,omitEmpty"`
	Origin    string      `json:"origin,omitEmpty"`
	Meanings  []Meaning   `json:"meanings,omitEmpty"`
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
	Antonyms []string `json:"antonyms,omitEmpty"`
}

type dictionaryResponse struct {
	Response []*Entry
}

type CommandDefinition struct {
	Word string
	POS  string
	Def  string
}

type DefError struct {
	Title      string `json:"title"`
	Message    string `json:"message"`
	Resolution string `json:"resolution"`
}
