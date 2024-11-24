package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type PropertiesConstant struct {
	key   string
	value interface{}
}

var propertiesConstants = make(map[string]PropertiesConstant)

func newProperty(key string, value interface{}) (PropertiesConstant, error) {
	return PropertiesConstant{key, value}, nil
}

func addProperty(key string, value interface{}) error {
	if _, exists := propertiesConstants[key]; exists {
		return errors.New(fmt.Sprint("Key ", key, " already exists."))
	}
	val, err := newProperty(key, value)
	if err != nil {
		return err
	}
	propertiesConstants[key] = val
	return nil
}

func GetProperty(key string) (interface{}, error) {
	value, exists := propertiesConstants[key]
	if !exists {
		return nil, errors.New(fmt.Sprint("Key ", key, " doesn't exist."))
	}
	return value.value, nil
}

const (
	DatasourceEngineName   = "datasource.engine"
	DatasourceUsername     = "datasource.username"
	DatasourcePassword     = "datasource.password"
	DatasourceUrl          = "datasource.url"
	DatasourceDatabaseName = "datasource.database.name"
)

func LoadProperties() error {
	file, err := os.Open("application.properties")
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		er := file.Close()
		if er != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid line: %s", line)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		err := addProperty(key, value)
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
