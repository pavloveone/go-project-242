package code

import (
	"fmt"
	"os"
)

func GetSize(path string) (string, error) {
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
			size += info.Size()
		}
	}
	return fmt.Sprintf("%d\t%s", size, path), nil
}
