package main

import (
	"GopherNetwork/internal/db"
	"GopherNetwork/internal/env"
	"GopherNetwork/internal/store"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Incorrect loading of dotenv")
	}
	addr := env.GetString("DB_ADDR", "postgres://postgres:1234@localhost/gophernetwork?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	store := store.NewStorage(conn)
	db.Seed(store)
}
