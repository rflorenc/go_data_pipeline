package stream

import (
	"fmt"
	"strconv"
	"strings"
)

/**
ConfigError represents a configuration error message.
*/
type ConfigError struct {
	error
	message string
}

/**
Config interface provides the means to access configuration
*/
type Config interface {
	GetString(key string) string
	GetInt(key string) int
	ToString() string
}

type PropertiesConfig struct {
	Config
	properties map[string]string
}

func LoadProperties(filename string) (Config, error) {
	lines, err := LoadTextFile(filename)
	if err != nil {
		return nil, nil
	}

	var raw = make(map[string]string)

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}

		pair := strings.Split(line, "=")

		if len(pair) != 2 {
			return nil, &ConfigError{message: fmt.Sprintf("invalid property format: %s", pair)}
		}

		key := strings.TrimSpace(pair[0])
		value := strings.TrimSpace(pair[1])

		raw[key] = value
	}

	config := &PropertiesConfig{properties: raw}
	return config, nil
}

func (config *PropertiesConfig) GetInt(key string) int {
	value, _ := strconv.Atoi(config.properties[key])
	return value
}

func (config *PropertiesConfig) ToString() string {
	return fmt.Sprintf("%s", config.properties)
}

/**
Returns the error message.
*/
func (configError *ConfigError) Error() string {
	return configError.message
}

func NewPropertiesConfig() Config {
	return &PropertiesConfig{properties: make(map[string]string)}
}
