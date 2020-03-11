package main

import (
	"database/sql"
	"fmt"
	phonedb "gophercises/ex8/db"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "gorniak"
	password = "linuxpl"
	dbname   = "gophercises"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s  sslmode=disable", host, port, user)
	must(phonedb.Reset("postgres", psqlInfo, dbname))

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	must(phonedb.Migrate("postgres", psqlInfo))

	db, err := phonedb.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	err = db.Seed()
	must(err)
	/*
		phones, err := allPhones(db)
		must(err)
		for _, p := range phones {
			fmt.Printf("Working on...%+v\n", p)
			number := normalize(p.number)
			if number != p.number {
				fmt.Println("Updating or removing...", number)
				existing, err := findPhone(db, number)
				must(err)
				if existing != nil {
					must(deletePhone(db, p.id))
				} else {
					p.number = number
					must(updatePhone(db, p))
				}
			} else {
				fmt.Println("No changes required")
			}
		} */
}

type phone struct {
	id     int
	number string
}

func deletePhone(db *sql.DB, id int) error {
	statement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.Exec(statement, id)
	return err
}

func updatePhone(db *sql.DB, p phone) error {
	statement := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.Exec(statement, p.id, p.number)
	return err
}

func allPhones(db *sql.DB) ([]phone, error) {
	rows, err := db.Query("SELECT id, value FROM phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []phone
	for rows.Next() {
		var p phone
		if err := rows.Scan(&p.id, &p.number); err != nil {
			return nil, err
		}
		result = append(result, p)
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func findPhone(db *sql.DB, number string) (*phone, error) {
	var p phone
	row := db.QueryRow("SELECT * FROM phone_numbers WHERE value=$1", number)
	err := row.Scan(&p.id, &p.number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err

	}
	return &p, nil
}
func getPhone(db *sql.DB, id int) (string, error) {
	var number string
	err := db.QueryRow("SELECT value FROM phone_numbers WHERE id=$1", id).Scan(&number)
	if err != nil {
		return "", nil
	}
	return number, nil
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil

}

func normalize(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

/* func normalize(phone string) string {
	var buf bytes.Buffer
	for _, digit := range phone {
		if digit < '0' || digit > '9' {
			continue
		}
		buf.WriteRune(digit)
	}
	return buf.String()
}
*/
