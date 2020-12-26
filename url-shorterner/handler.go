package urlshort

import (
	"net/http"
	"fmt"
	yaml "gopkg.in/yaml.v2"
	json "encoding/json"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		dest, ok := pathsToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w,r)
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

type Path struct {
	Path string
	Url string
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYAML, err := parseYAML(yml)
	if err != nil {
		return nil,err
	}
	pathMap := buildMap(parsedYAML) 
	fmt.Print(pathMap)
	return MapHandler(pathMap, fallback), nil
}

func buildMap(parsedData []Path) map[string]string {
	m := make(map[string]string, len(parsedData))
	for _, each := range parsedData {
		m[each.Path] = each.Url
	} 
	return m
}

func parseYAML(yml []byte) ([]Path ,error) {
	path := make([]Path, 0)
	err := yaml.Unmarshal(yml, &path)
	if err != nil {
		return nil, err
	}
	return path,nil
}


func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc,error) {
	parsedYAML, err := parseJSON(json)
	if err != nil {
		return nil,err
	}
	pathMap := buildMap(parsedYAML) 
	fmt.Print(pathMap)
	return MapHandler(pathMap, fallback), nil
}

func parseJSON(data []byte) ([]Path, error){
	path := make([]Path, 0)
	err := json.Unmarshal(data, &path)
	if err != nil {
		return nil, err
	}
	return path,nil
}