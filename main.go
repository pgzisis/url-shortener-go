package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pgzisis/url-shortener-go/urlshort"
)

func main() {
	yamlFileName, jsonFileName := getFileNames()
	yaml := readFile(yamlFileName)
	json := readFile(jsonFileName)

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	jsonHandler, err := urlshort.JSONHandler(json, yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func getFileNames() (*string, *string) {
	yamlFileName := flag.String("yaml", "pathsToUrls.yaml", "yaml file to read")
	jsonFileName := flag.String("json", "pathsToUrls.json", "json file to read")
	flag.Parse()

	return yamlFileName, jsonFileName
}

func readFile(fileName *string) []byte {
	file, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	return file
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
