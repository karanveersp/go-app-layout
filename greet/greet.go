package greet

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// Config struct for loading greetings
type Config struct {
	SpanishGreeting string `yaml:"SpanishGreeting"`
	EnglishGreeting string `yaml:"EnglishGreeting"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config struct
	config := &Config{}
	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	d := yaml.NewDecoder(file)

	// Start YAML decoding
	if err := d.Decode(&config); err != nil {
		return nil, fmt.Errorf("trouble unmarshalling config: %s", err)
	}

	return config, nil
}

func SayHello(config *Config, lang, name string) {
	if lang == "spanish" {
		fmt.Println(config.SpanishGreeting, name)
	} else {
		fmt.Println(config.EnglishGreeting, name)
	}
}
