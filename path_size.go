package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetPathSize вычисляет размер файла или директории по указанному пути.
//
// Если path указывает на файл, возвращается его размер.
// Если размер указывает на директорию, возвращается размер суммы файлов, расположенных внутри нее.
//
// Параметры:
//
//   - path - путь к файлу или директории.
//
//   - human - если true, результат возвращается в человекочитаемом формате (KB, MB, GB и тд),
//     иначе в байтах.
//
//   - all - если true, то учитываются скрытые файлы и директории (которые начинаются с ".").
//
//   - recursive - если true, то директория обходится рекурсивно.
//     Cчитается размер всех вложенных файлов
//
// Возвращаемые значения:
//
//   - строка с размером (например "123B", "1.5MB")
//
//   - ошибка, если путь не существует или призошла ошибка чтения
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
