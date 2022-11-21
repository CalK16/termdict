package core

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func db_path() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	os.Mkdir(filepath.Join(dirname, ".termdict"), 0755)
	return filepath.Join(dirname, ".termdict", "termdict.db")
}

func init() {
	// Create database if not exists
	db, err := sql.Open("sqlite3", db_path())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create table if not exists
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS words (
		spell TEXT,
		phonetic TEXT,
		pos TEXT,
		def TEXT,
		exp TEXT
	)	
`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}

func save(word Word) {
	db, err := sql.Open("sqlite3", db_path())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO words(spell, phonetic, pos, def, exp) values(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	for _, meaning := range word.Meanings {
		for _, definition := range meaning.Definitions {
			_, err = stmt.Exec(
				word.Spell,
				word.Phonetic,
				meaning.PartOfSpeech,
				definition.Definition,
				definition.Example,
			)
			if err != nil {
				panic(err)
			}
		}
	}
}

func read(spell string) (bool, Word) {
	db, err := sql.Open("sqlite3", db_path())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM words WHERE spell = ?", spell)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var meanings []Meaning
	var definitions []Definition
	var word Word
	var phonetic string

	found := false
	for rows.Next() {
		found = true
		var pos, def, exp string
		err = rows.Scan(&spell, &phonetic, &pos, &def, &exp)
		if err != nil {
			panic(err)
		}
		definition := Definition{Definition: def, Example: exp}
		definitions = append(definitions, definition)
		meaning := Meaning{PartOfSpeech: pos, Definitions: definitions}
		meanings = append(meanings, meaning)
	}
	word = Word{Spell: spell, Phonetic: phonetic, Meanings: meanings}
	return found, word
}
