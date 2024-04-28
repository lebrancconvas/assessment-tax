package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lebrancconvas/assessment-tax/db"
	"github.com/lebrancconvas/assessment-tax/server"
)

func init() {
	// os.Setenv("PORT", "8080")
	// os.Setenv("DBHOST", "localhost")
	// os.Setenv("DBPORT", "5432")
	// os.Setenv("DBUSER", "postgres")
	// os.Setenv("DBPASSWORD", "postgres")
	// os.Setenv("DBNAME", "ktaxes")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load .env")
	}

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()

	// Database Connect.
	db.Init()

	// Server Start
	server.Start()
}
