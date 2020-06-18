package provisioner

import (
	"encoding/json"
	"os"
)

func readJsonFile(path string) (map[string]interface{}, error) {
	var result map[string]interface{}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&result)
	return result, err
}
