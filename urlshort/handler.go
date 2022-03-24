package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathsToUrls []pathToUrl

	err := yaml.Unmarshal(yml, &pathsToUrls)
	if err != nil {
		return nil, err
	}

	return handler(pathsToUrls, fallback)
}

func JSONHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathsToUrls []pathToUrl

	err := json.Unmarshal(data, &pathsToUrls)
	if err != nil {
		return nil, err
	}

	return handler(pathsToUrls, fallback)
}

func handler(pathsToUrls []pathToUrl, fallback http.Handler) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		for _, pathToUrl := range pathsToUrls {
			if path == pathToUrl.Path {
				http.Redirect(w, r, pathToUrl.Url, http.StatusFound)
			}
		}

		fallback.ServeHTTP(w, r)
	}, nil
}

type pathToUrl struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}
