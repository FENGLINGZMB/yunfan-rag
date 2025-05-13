package config

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config, err := NewConfig("")
	fmt.Println(err)
	jsonData, err := json.Marshal(config)
	fmt.Println(err)
	fmt.Println(string(jsonData))
}

