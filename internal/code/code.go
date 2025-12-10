package code

import (
	"fmt"
	"os"
	"strings"
)

func GetSize(path string, human, all bool) (string, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", err
	}
	var size int64
	if !fileInfo.IsDir() {
		size = fileInfo.Size()
	} else {
		entries, err := os.ReadDir(path)
		if err != nil {
			return "", err
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			info, err := entry.Info()
			if err != nil {
				return "", err
			}
			if isHidden(info.Name(), all) {
				continue
			}
			size += info.Size()
		}
	}
	formated, err := FormatSize(size, human)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\t%s", formated, path), nil
}

func FormatSize(size int64, human bool) (string, error) {
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
