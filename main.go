package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

//go:embed languages.yml
var embeddedYAML embed.FS

const (
	yamlURL               = "https://raw.githubusercontent.com/github/linguist/master/lib/linguist/languages.yml"
	extensionToLangFile   = "extension_to_language.json"
	langToExtensionFile   = "language_to_extension.json"
	detailedExtToLangFile = "detailed_extension_to_language.json"
	detailedLangToExtFile = "detailed_language_to_extension.json"
)

type LanguageData struct {
	Type               string   `yaml:"type" json:"type"`
	Color              string   `yaml:"color" json:"color"`
	Extensions         []string `yaml:"extensions,omitempty" json:"extensions,omitempty"`
	Filenames          []string `yaml:"filenames,omitempty" json:"filenames,omitempty"`
	AceMode            string   `yaml:"ace_mode" json:"ace_mode"`
	CodemirrorMode     string   `yaml:"codemirror_mode,omitempty" json:"codemirror_mode,omitempty"`
	CodemirrorMimeType string   `yaml:"codemirror_mime_type,omitempty" json:"codemirror_mime_type,omitempty"`
	LanguageId         int      `yaml:"language_id" json:"language_id"`
}

func main() {
	// Define flags
	var detailed bool
	flag.BoolVar(&detailed, "d", false, "Generate detailed language maps")
	flag.BoolVar(&detailed, "detailed", false, "Generate detailed language maps")

	// Customize usage output
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  -d, --detailed\tGenerate detailed language maps\n")
	}

	flag.Parse()

	// Fetch YAML data
	data, err := fetchYAML()
	if err != nil {
		log.Fatalf("Error obtaining YAML data: %v", err)
	}

	// Parse the YAML data
	var languagesMap map[string]LanguageData
	err = yaml.Unmarshal(data, &languagesMap)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	if detailed {
		createDetailedMaps(languagesMap)
	} else {
		createSimpleMaps(languagesMap)
	}
}

func fetchYAML() ([]byte, error) {
	resp, err := http.Get(yamlURL)
	if err != nil {
		log.Printf("Error fetching YAML from URL: %v", err)
		log.Println("Falling back to embedded YAML file...")
		return embeddedYAML.ReadFile("languages.yml")
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func createSimpleMaps(languagesMap map[string]LanguageData) {
	extensionToLang := make(map[string][]string)
	langToExtension := make(map[string][]string)

	for langName, langData := range languagesMap {
		for _, ext := range langData.Extensions {
			if !strings.HasPrefix(ext, ".") {
				ext = "." + ext
			}
			extensionToLang[ext] = append(extensionToLang[ext], langName)
			langToExtension[langName] = append(langToExtension[langName], ext)
		}
		for _, filename := range langData.Filenames {
			extensionToLang[filename] = append(extensionToLang[filename], langName)
			langToExtension[langName] = append(langToExtension[langName], filename)
		}
	}

	createJSONFile(extensionToLang, extensionToLangFile)
	createJSONFile(langToExtension, langToExtensionFile)

	fmt.Printf("Simple mapping JSON files created: %s and %s\n", extensionToLangFile, langToExtensionFile)
}

func createDetailedMaps(languagesMap map[string]LanguageData) {
	detailedExtToLang := make(map[string]map[string]LanguageData)
	detailedLangToExt := make(map[string]LanguageData)

	for langName, langData := range languagesMap {
		detailedLangToExt[langName] = langData

		for _, ext := range langData.Extensions {
			if !strings.HasPrefix(ext, ".") {
				ext = "." + ext
			}
			if _, exists := detailedExtToLang[ext]; !exists {
				detailedExtToLang[ext] = make(map[string]LanguageData)
			}
			detailedExtToLang[ext][langName] = langData
		}

		for _, filename := range langData.Filenames {
			if _, exists := detailedExtToLang[filename]; !exists {
				detailedExtToLang[filename] = make(map[string]LanguageData)
			}
			detailedExtToLang[filename][langName] = langData
		}
	}

	createJSONFile(detailedExtToLang, detailedExtToLangFile)
	createJSONFile(detailedLangToExt, detailedLangToExtFile)

	fmt.Printf("Detailed mapping JSON files created: %s and %s\n", detailedExtToLangFile, detailedLangToExtFile)
}

func createJSONFile(data interface{}, filename string) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error creating JSON data for %s: %v", filename, err)
	}

	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing JSON file %s: %v", filename, err)
	}
}
