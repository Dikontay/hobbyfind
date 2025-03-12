package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

func InitConfigs(filePath string, configs interface{}) error {
	if reflect.TypeOf(configs).Kind() != reflect.Ptr {
		return fmt.Errorf("filePath must be a string")
	}

	ext := filepath.Ext(filePath)

	if ext != ".json" {
		return fmt.Errorf("configs file must have .json extension")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read data err: %v", err)
	}

	err = json.Unmarshal(data, configs)
	if err != nil {
		return fmt.Errorf("failed to unmarshall JSON: %v", err)
	}
	return nil

}
