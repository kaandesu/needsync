package needs

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func LoadAll(dir string) ([]*Need, error) {
	var needs []*Need

	files, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}

		var n Need
		if err := json.Unmarshal(data, &n); err != nil {
			return nil, err
		}

		needs = append(needs, &n)
	}

	return needs, nil
}

func Save(path string, n *Need) error {
	data, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
