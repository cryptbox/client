package main

import (
	"flag"
	"net/http"
	"os"
	"path/filepath"
)

var rootPath = flag.String("root", ".", "Root folder to share")

func StartWebServer() {
	flag.Parse()

	info, err := os.Stat(*rootPath)
	if os.IsNotExist(err) || !info.IsDir() {
		log.Fatal("Invalid root path:", *rootPath)
	}

	path, _ := filepath.Abs(*rootPath)
	log.Info("Serving:", path)

	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(*rootPath))).Error())
}
