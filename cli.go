package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func today() string {
	return time.Now().Format("2006-01-02")
}

type CLI struct {
	DBFile string
	db     *DB
}

func (cli *CLI) openDB() error {
	cli.db = NewDB(cli.DBFile)
	err := cli.db.Load()
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("failed to open db (%v)", err)
	}
	fmt.Printf("info: failed to open db, creating new one\n")
	err = cli.db.Save()
	if err != nil {
		return fmt.Errorf("failed to write db", err)
	}
	return nil
}

func (cli *CLI) Run(args []string) int {
	err := cli.openDB()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return 1
	}
	if len(args) == 0 {
		return cli.ListAll()
	}
	if len(args) == 1 && validDate(args[0]) {
		return cli.ListDate(args[0])
	}
	date := today()
	textArgs := args
	if validDate(args[0]) {
		date = args[0]
		textArgs = args[1:]
	}
	return cli.AddEntry(date, strings.Join(textArgs, " "))
}

func (cli *CLI) ListAll() int {
	for _, date := range cli.db.SortedDates() {
		cli.listDate(date)
	}
	return 0
}

func (cli *CLI) listDate(date string) {
	entries := cli.db.EntriesForDate(date)
	if len(entries) < 1 {
		return
	}
	fmt.Printf("%s\n==========\n", date)
	for _, entry := range entries {
		fmt.Printf("  * %s\n", entry)
	}
	fmt.Printf("\n")
}

func (cli *CLI) ListDate(date string) int {
	cli.listDate(date)
	return 0
}

func (cli *CLI) AddEntry(date, text string) int {
	cli.db.AddEntry(date, text)
	err := cli.db.Save()
	if err != nil {
		fmt.Printf("error: failed to save db (%v)\n", err)
		return 1
	}
	return 0
}
