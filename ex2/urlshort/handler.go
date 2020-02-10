package urlshort

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

	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.URL.Path
		normalURL := pathsToUrls[shortURL]
		if normalURL == "" {
			fallback.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, normalURL, http.StatusFound)
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
	YamlRecords, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(YamlRecords)
	return MapHandler(pathMap, fallback), nil
}

type urlRecord struct {
	URL  string `yaml:"url"`
	Path string `yaml:"path"`
}

func parseYAML(in []byte) ([]urlRecord, error) {

	var url []urlRecord
	err := yaml.Unmarshal(in, &url)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func buildMap(YamlRecords []urlRecord) map[string]string {
	var result = make(map[string]string, len(YamlRecords))
	for _, record := range YamlRecords {
		result[record.Path] = record.URL
	}
	return result
}
