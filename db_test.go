package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestWrite(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { os.RemoveAll(tmpdir) }()

	dbFile := filepath.Join(tmpdir, "db")
	db := NewDB(dbFile)
	db.AddEntry("2015-05-05", "Started implementing dalo +1")
	db.AddEntry("2015-05-05", "Finished implementing dalo +2")
	err = db.Save()
	if err != nil {
		t.Fatal(err)
	}
	db2 := NewDB(dbFile)
	err = db2.Load()
	if err != nil {
		t.Fatal(err)
	}
	entries := db2.EntriesForDate("2015-05-05")
	if len(entries) != 2 {
		t.Fatal("wrong number of entries")
	}
}
