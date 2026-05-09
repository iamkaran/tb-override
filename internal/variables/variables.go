// Package variables contains helper functions for fetching items from the list of css attributes
package variables

import (
	"encoding/json"
	"os"
)

// VariablesJSON holds the JSON form of all the css attributes
type VariablesJSON struct {
	Data map[string]map[string]map[string]string
}

// FetchCategories returns the list of categories
func (v *VariablesJSON) FetchCategories() []string {
	keys := make([]string, 0, len(v.Data))

	for k := range v.Data {
		keys = append(keys, k)
	}

	return keys
}

// FetchItems returns vars belonging to a category
func (v *VariablesJSON) FetchItems(category string) map[string]map[string]string {
	return v.Data[category]
}

// FetchVariables returns the entire JSON map (use only when neccessary)
func (v *VariablesJSON) FetchVariables() map[string]map[string]map[string]string {
	return v.Data
}

func (v *VariablesJSON) LoadMap(filename string) (*VariablesJSON, error) {
	data, _ := os.ReadFile(filename)

	jsonData := make(map[string]map[string]map[string]string)
	err := json.Unmarshal(data, &jsonData)

	if err != nil {
		return nil, err
	}

	return &VariablesJSON{
		Data: jsonData,
	}, nil
}
