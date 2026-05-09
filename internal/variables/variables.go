// Package variables contains helper functions for fetching items from the list of css attributes
package variables

import (
	"encoding/json"
	"os"
)

// VariablesJSON holds the JSON form of all the css attributes
type VariablesJSON struct {
	Data map[string][]Variable
}

type Variable struct {
	Name        string
	Default     string
	Type        string
	Description string
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
func (v *VariablesJSON) FetchItems(category string) []Variable {
	return v.Data[category]
}

// FetchVariables returns the entire JSON map (use only when neccessary)
func (v *VariablesJSON) FetchVariables() map[string][]Variable {
	return v.Data
}

func LoadMap(filename string) (*VariablesJSON, error) {
	data, _ := os.ReadFile(filename)

	raw := make(map[string]map[string]map[string]string)
	err := json.Unmarshal(data, &raw)

	if err != nil {
		return nil, err
	}

	normalized := make(map[string][]Variable)

	for category, vars := range raw {
		for name, meta := range vars {
			normalized[category] = append(normalized[category], Variable{
				Name:        name,
				Default:     meta["default"],
				Type:        meta["type"],
				Description: meta["description"],
			})
		}
	}

	return &VariablesJSON{
		Data: normalized,
	}, nil
}
