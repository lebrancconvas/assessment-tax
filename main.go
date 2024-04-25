package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lebrancconvas/assessment-tax/db"
	"github.com/lebrancconvas/assessment-tax/server"
)

func init() {
	os.Setenv("PORT", "8080")
	os.Setenv("DBHOST", "localhost")
	os.Setenv("DBPORT", "5432")
	os.Setenv("DBUSER", "postgres")
	os.Setenv("DBPASSWORD", "postgres")
	os.Setenv("DBNAME", "ktaxes")
}

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()

	// Database Start
	db.Init()

	// Server Start
	server.NewRouter()
}
