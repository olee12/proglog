package main

import (
	"log"

	"github.com/olee12/proglog/internal/server"
)

func main() {
	svr := server.NewHTTPServer(":8080")
	log.Fatal(svr.ListenAndServe())
}
