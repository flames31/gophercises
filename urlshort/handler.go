package urlshort

import (
	"fmt"
	"net/http"

	"github.com/go-yaml/yaml"
)

type urlPathType struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if destURL, ok := pathsToUrls[url]; ok {
			http.Redirect(w, r, destURL, http.StatusFound)
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
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(data []byte) ([]urlPathType, error) {
	urlPaths := []urlPathType{}
	err := yaml.Unmarshal(data, &urlPaths)
	if err != nil {
		return []urlPathType{}, fmt.Errorf("ERROR : %w", err)
	}
	return urlPaths, nil
}

func buildMap(urls []urlPathType) map[string]string {
	urlMap := make(map[string]string)
	for _, url := range urls {
		urlMap[url.Path] = url.Url
	}

	return urlMap
}
