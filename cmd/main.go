package main

import (
	"context"
	"log"

	"github.com/eugbyte/sqlc_goose_tutorial/internal/db"
	"github.com/eugbyte/sqlc_goose_tutorial/internal/db/codegen"
)

func run() error {
	ctx := context.Background()

	conn := db.OpenDbConnection(ctx, "user=postgres password=postgres dbname=book_demo sslmode=disable")
	defer conn.Close(ctx)

	var queries *codegen.Queries = codegen.New(conn)

	// list all books
	books, err := queries.ListBooks(ctx, codegen.ListBooksParams{})
	if err != nil {
		return err
	}
	log.Println(books)

	// create a book
	insertedBook, err := queries.CreateBook(ctx, codegen.CreateBookParams{
		Title:  "The Go Programming Language",
		Author: "Alan Donovan & Brian Kernighan",
	})
	if err != nil {
		return err
	}
	log.Println(insertedBook)

	// get the book we just inserted
	fetchedBook, err := queries.GetBook(ctx, insertedBook.ID)
	if err != nil {
		return err
	}
	log.Println(fetchedBook)

	// list books filtered by author; nil fields are not filtered on
	kernighanBooks, err := queries.ListBooks(ctx, codegen.ListBooksParams{
		Author: new("kernighan"),
	})
	if err != nil {
		return err
	}
	log.Println(kernighanBooks)

	// update the book
	err = queries.UpdateBook(ctx, codegen.UpdateBookParams{
		ID:     insertedBook.ID,
		Title:  "The C Programming Language",
		Author: "Brian Kernighan & Dennis Ritchie",
	})
	if err != nil {
		return err
	}

	// delete the book
	return queries.DeleteBook(ctx, insertedBook.ID)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
