package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

var TomlMap = make(map[string]map[string]string)

func TomlInit() {
	folderPath := "./settings"

	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatalf("❌ Failed to read settings folder: %v", err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".toml") {
			continue
		}

		filePath := filepath.Join(folderPath, file.Name())
		innerMap := make(map[string]string)

		if _, err := toml.DecodeFile(filePath, &innerMap); err != nil {
			log.Fatalf("❌ Error loading TOML file %s: %v", file.Name(), err)
		}

		key := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		TomlMap[key] = innerMap
	}

	fmt.Println("✅ TOML configuration loaded successfully")
}

func GetTomlValue(fileName, key string) string {
	if valMap, ok := TomlMap[fileName]; ok {
		if val, exists := valMap[key]; exists {
			return val
		}
		log.Printf("⚠️ TOML key not found: [%s][%s]", fileName, key)
	}
	return ""
}
