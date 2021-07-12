package main

import (
	"log"
	"net/http"
)

func main() {
	//port := os.Getenv("PORT")

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Executando na porta 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}