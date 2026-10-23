package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/eugbyte/sqlc_goose_tutorial/internal/db"
	"github.com/eugbyte/sqlc_goose_tutorial/internal/db/codegen"
)

func run(ctx context.Context) error {
	conn := db.OpenDbConnection(ctx, "user=postgres password=postgres dbname=book_demo sslmode=disable")
	defer conn.Close(ctx)

	var repo codegen.Querier = codegen.New(conn)

	// start from an empty table on every run
	deleted, err := repo.DeleteAllBooks(ctx)
	if err != nil {
		return err
	}
	log.Printf("deleted %d books", deleted)

	// create a book
	insertedBook, err := repo.CreateBook(ctx, codegen.CreateBookParams{
		Title:  "The Go Programming Language",
		Author: "Alan Donovan & Brian Kernighan",
	})
	if err != nil {
		return err
	}
	log.Println("Created books...")
	log.Println(insertedBook)

	// get the book we just inserted
	fetchedBook, err := repo.GetBook(ctx, insertedBook.ID)
	if err != nil {
		return err
	}
	log.Println(fetchedBook)

	// list books filtered by author; nil fields are not filtered on
	kernighanBooks, err := repo.ListBooks(ctx, codegen.ListBooksParams{
		Author: new("kernighan"),
	})
	if err != nil {
		return err
	}
	log.Println("Listing books...")
	log.Println(kernighanBooks)

	// update the book
	err = repo.UpdateBook(ctx, codegen.UpdateBookParams{
		ID:     insertedBook.ID,
		Title:  "The C Programming Language",
		Author: "Brian Kernighan & Dennis Ritchie",
	})
	if err != nil {
		return err
	}
	kernighanBooks, err = repo.ListBooks(ctx, codegen.ListBooksParams{
		Author: new("kernighan"),
	})
	if err == nil {
		log.Println("Updated books...")
		log.Println(kernighanBooks)
	}

	return err
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()
}
