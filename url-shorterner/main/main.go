package main

import (
	"fmt"
	"net/http"
	"urlshort"
	"flag"
	"io/ioutil"
)

func main() {
	var ymlFileName = flag.String("yml", "handler.yml", "a yaml file in the format of path and url to add routes")
	flag.Parse()

	yaml, err := ioutil.ReadFile(*ymlFileName)
	if err != nil {
		panic(err)
	}
	fmt.Print(yaml) // temporary
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
// 	yaml := `
// - path: /urlshort
//   url: https://github.com/gophercises/urlshort
// - path: /urlshort-final
//   url: https://github.com/gophercises/urlshort/tree/solution
// `
	// yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	// if err != nil {
	// 	panic(err)
	// }

	// Build a JSON handler
	// https://blog.golang.org/json
	// package: https://golang.org/pkg/encoding/json/
	json := `[{"path": "/urlshort", "url": "https://github.com/gophercises/urlshort"}, {"path": "/urlshort-final", "url": "https://github.com/gophercises/urlshort/tree/solution"}]`

	jsonHandler, err := urlshort.JSONHandler([]byte(json), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
