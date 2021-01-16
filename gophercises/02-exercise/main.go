package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)
const defaultFile = "data.yaml"
func main() {
	mux := defaultMux()
	filePath := flag.String("file", defaultFile, "A yaml file with urls and paths")
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	flag.Parse()
	yaml, readError := ioutil.ReadFile(*filePath)
	if readError != nil {
		panic(readError)
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	yamlHandler, parseError := YAMLHandler(yaml, mapHandler)
	if parseError != nil {
		panic(parseError)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
