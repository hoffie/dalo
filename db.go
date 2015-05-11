package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

type DB struct {
	file    string
	entries map[string][]string
}

// warning: fails to catch 2015-02-31 etc.
var dateRegexp = regexp.MustCompile(`^(\d{4})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$`)

func validDate(date string) bool {
	return dateRegexp.MatchString(date)
}

// NewDB returns a new DB instance.
func NewDB(file string) *DB {
	file, _ = filepath.Abs(file)
	db := &DB{file: file, entries: make(map[string][]string)}
	return db
}

// AddEntry creates a new entry at the given date.
func (db *DB) AddEntry(date string, text string) {
	if _, ok := db.entries[date]; !ok {
		db.entries[date] = []string{text}
		return
	}
	db.entries[date] = append(db.entries[date], text)
}

// Save stores the internal state in the pre-configured file.
func (db *DB) Save() error {
	tmpfile := db.file + ".tmp"
	data, err := db.toBytes()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(tmpfile, data, 0600)
	if err != nil {
		return err
	}
	err = os.Rename(tmpfile, db.file)
	return err
}

// Load loads the pre-configured file into the internal state.
func (db *DB) Load() error {
	data, err := ioutil.ReadFile(db.file)
	if err != nil {
		return err
	}
	return db.fromBytes(data)
}

// toBytes returns the internal store's byte representation.
func (db *DB) toBytes() ([]byte, error) {
	return json.MarshalIndent(db.entries, "", "  ")
}

// fromBytes parses the given representation and sets it as the internal store.
func (db *DB) fromBytes(b []byte) error {
	entries := map[string][]string{}
	err := json.Unmarshal(b, &entries)
	if err != nil {
		return err
	}
	db.entries = entries
	return nil
}

// SortedDates returns the list of all dates, sorted from newest to oldest
func (db *DB) SortedDates() []string {
	dates := make([]string, len(db.entries))
	i := 0
	for key, _ := range db.entries {
		dates[i] = key
		i++
	}
	sort.Strings(dates)
	return dates
}

// EntriesForDate returns all entries which are stored for the given date
func (db *DB) EntriesForDate(date string) []string {
	entries, ok := db.entries[date]
	if !ok {
		return []string{}
	}
	return entries
}
