package main

import (
	"flag"
	"log"

	"github.com/hramcovdv/cm_info/api"
	"github.com/hramcovdv/cm_info/storage"
)

var (
	listenAddr string
	dataSource string
)

func init() {
	flag.StringVar(&listenAddr, "listen", ":5161", "HTTP service listen address")
	flag.StringVar(&dataSource, "database", "data.sqlite", "Database filename")
	flag.Parse()
}

func main() {
	sqlite, err := storage.NewSqliteStorage(dataSource)
	if err != nil {
		log.Fatalf("NewSqliteStorage() error: %s", err)
	}
	defer sqlite.Close()

	// memory := storage.NewMemoryStorage()

	server := api.NewServer(sqlite)

	log.Printf("Server start listen %s", listenAddr)
	log.Fatal(server.Start(listenAddr))
}
