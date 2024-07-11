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
	//	TODO: Implement this...
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		for p, u := range pathsToUrls {
			if p == req.URL.Path {
				_, err := w.Write([]byte(u))
				if err != nil {
					break
				}
			}
		}
		fallback.ServeHTTP(w, req)
	})
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
type PathMappings struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	shortenedPaths := make([]PathMappings, 0)
	err := yaml.Unmarshal(yml, shortenedPaths)
	if err != nil {
		return nil, err
	}
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		for _, sp := range shortenedPaths {
			if sp.Path == req.URL.Path {
				_, err := w.Write([]byte(sp.Url))
				if err != nil {
					break
				}
			}
		}
		fallback.ServeHTTP(w, req)
	}), nil
}
