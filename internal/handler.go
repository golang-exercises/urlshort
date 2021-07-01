package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-yaml/yaml"
)

type YAMLData []struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

type JSONData []struct {
	Path string `json:"path"`
	Url  string `json:"url"`
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	mux := http.NewServeMux()

	for key, value := range pathsToUrls {
		path := key
		url := value
		mux.HandleFunc(path, func(res http.ResponseWriter, req *http.Request) {
			fmt.Printf("Map handler: %s\n", path)
			http.Redirect(res, req, url, http.StatusSeeOther)
		})
	}

	mux.Handle("/", fallback)

	return mux.ServeHTTP
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var yamlMap YAMLData
	mux := http.NewServeMux()
	err := yaml.Unmarshal(yml, &yamlMap)

	for _, value := range yamlMap {
		urlshort := value
		mux.HandleFunc(urlshort.Path, func(res http.ResponseWriter, req *http.Request) {
			fmt.Printf("YAML handler: %s\n", urlshort.Path)
			http.Redirect(res, req, urlshort.Url, http.StatusSeeOther)
		})
	}

	mux.Handle("/", fallback)

	return mux.ServeHTTP, err
}

func JSONHandler(jsonBlob []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var jsonMap JSONData
	mux := http.NewServeMux()
	err := json.Unmarshal(jsonBlob, &jsonMap)

	for _, value := range jsonMap {
		urlshort := value
		mux.HandleFunc(urlshort.Path, func(res http.ResponseWriter, req *http.Request) {
			fmt.Printf("JSON handler: %s\n", urlshort.Path)
			http.Redirect(res, req, urlshort.Url, http.StatusSeeOther)
		})
	}

	mux.Handle("/", fallback)

	return mux.ServeHTTP, err
}
