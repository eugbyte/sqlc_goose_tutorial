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
	deletedCount, err := repo.DeleteAllBooks(ctx)
	if err != nil {
		return err
	}
	log.Printf("deleted %d books", deletedCount)

	// create a book
	book, err := repo.CreateBook(ctx, codegen.CreateBookParams{
		Title:  "The Go Programming Language",
		Author: "Alan Donovan & Brian Kernighan",
	})
	if err != nil {
		return err
	}
	log.Println("Created books...")
	log.Println(book)

	// get the book we just inserted
	book, err = repo.GetBook(ctx, book.ID)
	if err != nil {
		return err
	}
	log.Println(book)

	// list books filtered by author; nil fields are not filtered on
	books, err := repo.ListBooks(ctx, codegen.ListBooksParams{
		Author: new("kernighan"),
	})
	if err != nil {
		return err
	}
	log.Println("Listing books...")
	log.Println(books)

	// update the book
	err = repo.UpdateBook(ctx, codegen.UpdateBookParams{
		ID:     book.ID,
		Title:  "The C Programming Language",
		Author: "Brian Kernighan & Dennis Ritchie",
	})
	if err != nil {
		return err
	}
	books, err = repo.ListBooks(ctx, codegen.ListBooksParams{
		Author: new("kernighan"),
	})
	if err == nil {
		log.Println("Updated books...")
		log.Println(books)
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
