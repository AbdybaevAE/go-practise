package main

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if val, ok := pathsToUrls[url]; ok {
			http.Redirect(w, r, val, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

type Item struct {
	Path string
	Url string
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	items, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}
 	pathsToUrls := buildMap(items)
	return MapHandler(pathsToUrls, fallback), nil
}
func parseYaml(yml[] byte) ([]Item, error) {
	items := make([]Item, 0)
	err := yaml.Unmarshal(yml, &items)
	return items, err
}
func buildMap(items []Item) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, item := range items {
		pathsToUrls[item.Path] = item.Url
	}
	return pathsToUrls
}