package server

import "os"

func Start() {
	router := NewRouter()
	router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))
}
