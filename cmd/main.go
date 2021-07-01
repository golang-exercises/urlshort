package main

import (
	"fmt"
	"net/http"
	"strconv"
	mod "urlshort/internal"
)

func main() {
	mux := defaultMux()
	yaml, json, port := mod.ParseFlags()
	db, dbData := mod.HandleDbConnection()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := mod.MapHandler(pathsToUrls, mux)
	yamlHandler, yamlErr := mod.YAMLHandler(yaml, mapHandler)
	jsonHandler, jsonErr := mod.JSONHandler(json, yamlHandler)
	dbHandler := mod.MapHandler(dbData, jsonHandler)

	if yamlErr != nil || jsonErr != nil {
		panic("File parsing error")
	}

	fmt.Printf("Server started on port %d.\n", port)
	http.ListenAndServe(":"+strconv.Itoa(port), dbHandler)

	defer db.Close()
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultFallback)
	return mux
}

func defaultFallback(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Default fallback: No valid path was detected.\n")
	fmt.Fprintln(w, "No valid path was detected")
}
