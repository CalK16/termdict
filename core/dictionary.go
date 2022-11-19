package core

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

type Word struct {
	Spell    string    `json:"word"`
	Phonetic string    `json:"phonetic"`
	Meanings []Meaning `json:"meanings"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
}

type Definition struct {
	Definition string `json:"definition"`
	Example    string `json:"example"`
}

func capitalize(s string) string {
	// Capitalize the first letter of the string
	return strings.ToUpper(s[:1]) + s[1:]
}

func TerminalFormatPrint(word Word) {
	const templ = "\033[1;31m{{.Spell | ToUpper }}\033[0m\n" +
		"\033[43m{{.Phonetic}}\033[0m\n" +
		"----------------------------------\n" +
		"Meanings:\n" +
		"{{range .Meanings}}" +
		"\033[31m[{{.PartOfSpeech}}]\033[0m\n" +
		"{{range .Definitions}}" +
		"• {{.Definition}}\n" +
		"  \033[2m{{.Example}}\033[0m\n" +
		"{{end}}" +
		"{{end}}"
	report, err := template.New("report").Funcs(template.FuncMap{"ToUpper": capitalize}).Parse(templ)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = report.Execute(os.Stdout, word)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Query(spell string) *Word {
	found, word := read(spell)
	if found {
		log.Default().Println("Found in local database")
		return &word
	}
	word, err := Crawl(spell)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	save(word)
	return &word
}
