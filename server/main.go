package main

import (
	"database/sql"
	"github.com/alexflint/go-arg"
	"log"
	"sync"

	"mailinglist-grpc/grpcapi"
	"mailinglist-grpc/mdb"
)

var args struct {
	DbPath   string `arg:"env:MAILINGLIST_DB"`
	BindGrpc string `arg:"env:MAILINGLIST_BIND_GRPC"`
}

func main() {
	arg.MustParse(&args)

	if args.DbPath == "" {
		args.DbPath = "list.db"
	}
	if args.BindGrpc == "" {
		args.BindGrpc = ":8080"
	}

	log.Printf("using database '%v'\n", args.DbPath)
	db, err := sql.Open("sqlite3", args.DbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mdb.TryCreate(db)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		log.Printf("starting gRPC API server...\n")
		grpcapi.Serve(db, args.BindGrpc)
		wg.Done()
	}()

	wg.Wait()
}
