package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pgzisis/url-shortener/urlshort"
)

func main() {
	yamlFileName := getYamlFileName()
	yaml := readYaml(yamlFileName)

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
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func getYamlFileName() *string {
	yamlFileName := flag.String("yaml", "pathsToUrls.yaml", "name of the yaml file to read")
	flag.Parse()

	return yamlFileName
}

func readYaml(yamlFileName *string) []byte {
	yaml, err := ioutil.ReadFile(*yamlFileName)
	if err != nil {
		panic(err)
	}

	return yaml
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
