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
		return "", fmt.Errorf("invalid size: size must be >= 0")
	}

	if !human {
		return fmt.Sprintf("%dB", size), nil
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

	if size == 0 {
		return "0B", nil
	}

	var unitIndex int
	value := float64(size)

	for value >= 1024 && unitIndex < len(units)-1 {
		value /= 1024
		unitIndex++
	}

	if unitIndex == 0 {
		return fmt.Sprintf("%dB", size), nil
	}
	formatted := fmt.Sprintf("%.1f%s", value, units[unitIndex])
	if strings.HasSuffix(formatted, ".0") {
		formatted = formatted[:len(formatted)-2] + units[unitIndex]
	}

	return formatted, nil
}

func isHidden(name string, all bool) bool {
	return strings.HasPrefix(name, ".") && !all
}

func calcSize(path string, all, recursive bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if !info.IsDir() {
		if isHidden(info.Name(), all) {
			return 0, nil
		}
		return info.Size(), nil
	}
	var totalSize int64

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		entryName := entry.Name()
		if isHidden(entryName, all) {
			continue
		}

		fullPath := filepath.Join(path, entryName)

		if entry.IsDir() {
			if recursive {
				dirSize, err := calcSize(fullPath, all, recursive)
				if err != nil {
					return 0, err
				}
				totalSize += dirSize
			}
			continue
		}
		fileInfo, err := entry.Info()
		if err != nil {
			return 0, err
		}
		totalSize += fileInfo.Size()
	}

	return totalSize, nil
}
