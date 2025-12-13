package path_size

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, human, all, recursive bool) (string, error) {
	size, err := calcSize(path, all, recursive)
	if err != nil {
		return "", err
	}
	formated, err := formatSize(size, human)
	if err != nil {
		return "", err
	}
	return formated, nil
}

func formatSize(size int64, human bool) (string, error) {
	if size < 0 {
		return "", fmt.Errorf("size must be < 0")
	}
	defRes := fmt.Sprintf("%dB", size)
	if !human {
		return defRes, nil
	}
	type unit struct {
		name  string
		value int64
	}
	units := []unit{
		{"B", 1},
		{"KB", 1024},
		{"MB", 1024 * 1024},
		{"GB", 1024 * 1024 * 1024},
		{"TB", 1024 * 1024 * 1024 * 1024},
		{"PB", 1024 * 1024 * 1024 * 1024 * 1024},
		{"EB", 1024 * 1024 * 1024 * 1024 * 1024 * 1024},
	}
	for i := len(units) - 1; i >= 0; i-- {
		u := units[i]
		if size >= u.value {
			if u.name == "B" {
				return defRes, nil
			}
			val := float64(size) / float64(u.value)
			return fmt.Sprintf("%.1f%s", val, u.name), nil
		}
	}
	return "0B", nil
}

func isHidden(name string, all bool) bool {
	return strings.HasPrefix(name, ".") && !all
}

func calcSize(path string, all, recursive bool) (int64, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if !fileInfo.IsDir() {
		if isHidden(fileInfo.Name(), all) {
			return 0, nil
		}
		return fileInfo.Size(), nil
	}
	var size int64
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	for _, entry := range entries {
		name := entry.Name()
		if isHidden(name, all) {
			continue
		}
		fullPath := filepath.Join(path, name)
		if entry.IsDir() {
			if !recursive {
				continue
			}
			s, err := calcSize(fullPath, all, recursive)
			if err != nil {
				return 0, err
			}
			size += s
			continue
		}
		info, err := entry.Info()
		if err != nil {
			return 0, err
		}
		size += info.Size()
	}
	return size, nil
}
