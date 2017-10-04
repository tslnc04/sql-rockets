package rockets

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	Host     string `json:"host"`
	DBname   string `json:"dbname"`
}

// LoadFile loads a file from the name into a byte slice
func LoadFile(filename string) []byte {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return file
}

// LoadConfigFromFile takes a byte slice of a file and generates a config object
func LoadConfigFromFile(file []byte) *Config {
	output := new(Config)
	err := json.Unmarshal(file, &output)

	if err != nil {
		panic(err)
	}

	return output
}
