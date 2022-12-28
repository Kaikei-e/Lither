package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed build/*
//go:embed build/index.html
//go:embed build/_app/*

var publicFS embed.FS

func main() {
	public, err := fs.Sub(publicFS, "build")
	if err != nil {
		panic(err)
	}
	http.Handle("/", http.FileServer((http.FS(public))))
	log.Fatal(http.ListenAndServe(":9000", nil))

}
