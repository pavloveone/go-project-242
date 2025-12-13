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
	return formatSize(size, human), nil
}

func formatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	if size == 0 {
		return "0B"
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	unitIndex := 0
	value := float64(size)

	for value >= 1024 && unitIndex < len(units)-1 {
		value /= 1024
		unitIndex++
	}

	if unitIndex == 0 {
		return fmt.Sprintf("%dB", size)
	}

	return fmt.Sprintf("%.1f%s", value, units[unitIndex])
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

	err = filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if currentPath == path {
			return nil
		}

		baseName := filepath.Base(currentPath)
		if isHidden(baseName, all) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if !recursive {
			rel, err := filepath.Rel(path, currentPath)
			if err != nil {
				return err
			}

			if strings.Contains(rel, string(filepath.Separator)) {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		if !info.IsDir() {
			totalSize += info.Size()
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return totalSize, nil
}
