package model

import (
	"database/sql"
)

var DB *sql.DB

func SaveBook(book *Book) error {
	_, err := DB.Exec("INSERT INTO books title VALUES ?", book.Title)
	return err
}

func GetBooksList() ([]Book, error) {
	rows, err := DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
