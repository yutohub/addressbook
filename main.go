package main

import (
	"addressbook/controller"
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var (
	mode string
)

func init() {
	flag.StringVar(&mode, "mode", "show", "Select which mode")
}

func main() {
	flag.Parse()
	db, err := sql.Open("sqlite", "addressbook.db")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	if err := controller.CreateTable(db); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	if err := controller.Run(mode, db); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
