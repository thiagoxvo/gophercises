package urlshort

import (
	"fmt"
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if destination, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, destination, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
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
	pathToUrls := make([]foo, 5)
	err := yaml.Unmarshal(yml, &pathToUrls)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Path: %v URL: %v", pathToUrls[0].Path, pathToUrls[0].URL)

	return nil, nil
}

type foo struct {
	Path string `yml:"path"`
	URL  string `yml:"url"`
}
