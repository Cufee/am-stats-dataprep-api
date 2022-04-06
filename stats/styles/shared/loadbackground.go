package shared

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func LoadBackground(fileName string) ([]byte, error) {
	files, err := os.ReadDir("assets")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.Split(file.Name(), ".")[0] == fileName {
			f, err := os.ReadFile(filepath.Join("assets", file.Name()))
			if err != nil {
				return nil, err
			}
			return f, nil
		}
	}

	return nil, fmt.Errorf("background file not found")
}
