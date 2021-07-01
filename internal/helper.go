package internal

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

func ParseFlags() ([]byte, []byte, int) {
	var yamlPath string
	var jsonPath string
	var port int
	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	Root := filepath.Join(filepath.Dir(b), "..")

	defaultYAML := filepath.Join(Root, "assets/test.yaml")
	defaultJSON := filepath.Join(Root, "assets/test.json")

	flag.StringVar(&yamlPath, "yaml", defaultYAML, "Yaml file for defining path and url")
	flag.StringVar(&jsonPath, "json", defaultJSON, "JSON file for defining path and url")
	flag.IntVar(&port, "port", 8080, "Port of the server")
	flag.Parse()

	yamlFile, yamlErr := ioutil.ReadFile(yamlPath)
	jsonFile, jsonErr := ioutil.ReadFile(jsonPath)

	if yamlErr != nil || jsonErr != nil {
		panic("Flag error")
	}

	return yamlFile, jsonFile, port
}
