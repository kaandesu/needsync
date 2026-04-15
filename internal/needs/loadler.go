package needs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func LoadAll(dir string) ([]*Need, error) {
	var needs []*Need

	fmt.Printf("[DEBUG] LoadAll: searching in %s\n", dir)
	files, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		fmt.Printf("[DEBUG] LoadAll: glob error: %v\n", err)
		return nil, err
	}

	fmt.Printf("[DEBUG] LoadAll: found %d .json files\n", len(files))
	for _, f := range files {
		fmt.Printf("[DEBUG] LoadAll: processing file: %s\n", f)
		data, err := os.ReadFile(f)
		if err != nil {
			fmt.Printf("[DEBUG] LoadAll: read error on %s: %v\n", f, err)
			return nil, err
		}

		var n Need
		if err := json.Unmarshal(data, &n); err != nil {
			fmt.Printf("[DEBUG] LoadAll: unmarshal error on %s: %v\n", f, err)
			return nil, err
		}

		needs = append(needs, &n)
	}

	fmt.Printf("[DEBUG] LoadAll: returning %d needs\n", len(needs))
	return needs, nil
}

func Save(path string, n *Need) error {
	data, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
