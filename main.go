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
//go:embed build/_app/*/*.js
//go:embed build/_app/immutable/chunks/*.js 
//go:embed build/_app/immutable/components/pages/*.js 
//go:embed static/output.css

var publicFS embed.FS

func main() {
	static, err := fs.Sub(publicFS, "static")
	if err != nil {
		panic(err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(static))))

	public, err := fs.Sub(publicFS, "build")
	if err != nil {
		panic(err)
	}
	http.Handle("/", http.FileServer((http.FS(public))))
	log.Fatal(http.ListenAndServe(":9000", nil))

}
