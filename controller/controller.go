package controller

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Record struct {
	ID    int64
	Name  string
	Phone string
}

func CreateTable(db *sql.DB) error {
	const sql = `
	CREATE TABLE IF NOT EXISTS addressbook (
			id    INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name  TEXT NOT NULL,
			phone TEXT NOT NULL
	);`
	if _, err := db.Exec(sql); err != nil {
		return err
	}
	return nil
}

func ShowRecords(db *sql.DB) error {
	fmt.Println("----------------------------------------")
	rows, err := db.Query("SELECT * FROM addressbook")
	if err != nil {
		return err
	}
	for rows.Next() {
		var r Record
		if err := rows.Scan(&r.ID, &r.Name, &r.Phone); err != nil {
			return err
		}
		fmt.Printf("ID: %d\tName: %s\tTEL: %s\n", r.ID, r.Name, r.Phone)
	}
	fmt.Println("----------------------------------------")
	return nil
}

func InputRecord(db *sql.DB) error {
	var r Record

	fmt.Print("Name > ")
	fmt.Scan(&r.Name)

	fmt.Print("TEL > ")
	fmt.Scan(&r.Phone)

	const sql = "INSERT INTO addressbook(name, phone) values (?,?)"
	_, err := db.Exec(sql, r.Name, r.Phone)
	if err != nil {
		return err
	}
	return nil
}

func UpdateRecord(db *sql.DB) error {
	var r Record

	fmt.Print("ID > ")
	fmt.Scan(&r.ID)

	fmt.Print("Name > ")
	fmt.Scan(&r.Name)

	fmt.Print("TEL > ")
	fmt.Scan(&r.Phone)

	const sql = "UPDATE addressbook SET name = ?, phone=? WHERE id = ?"
	_, err := db.Exec(sql, r.Name, r.Phone, r.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRecord(db *sql.DB) error {
	var r Record

	fmt.Print("ID > ")
	fmt.Scan(&r.ID)

	const sql = "DELETE FROM addressbook WHERE id = ?"
	_, err := db.Exec(sql, r.ID)
	if err != nil {
		return err
	}
	return nil
}

func DropTable(db *sql.DB) error {
	ans := "N"
	fmt.Println("Are you sure you want to reset? (y/N)")
	fmt.Scan(&ans)
	if ans == "y" {
		const sql = "DROP TABLE addressbook"
		_, err := db.Exec(sql)
		if err != nil {
			return err
		}
	} else {
		return nil
	}
	return nil
}

func Run(mode string, db *sql.DB) error {
	switch mode {
	case "show":
		if err := ShowRecords(db); err != nil {
			return err
		}
	case "insert":
		if err := InputRecord(db); err != nil {
			return err
		}
	case "update":
		if err := UpdateRecord(db); err != nil {
			return err
		}
	case "delete":
		if err := DeleteRecord(db); err != nil {
			return err
		}
	case "reset":
		if err := DropTable(db); err != nil {
			return err
		}
	default:
		s := "-----------------------------------\n" +
			"The following modes are supported\n" +
			"-----------------------------------\n" +
			"show:\tShow all registrations\n" +
			"insert:\tRegister new address\n" +
			"update:\tUpdate existing address\n" +
			"delete:\tSelect and delete\n" +
			"reset:\tDelete addressbook\n" +
			"-----------------------------------\n"
		fmt.Printf("%s", s)
	}
	return nil
}
