package utils

import (
	"os"
	"path/filepath"
)

func CreateDir(filePath string) error {
	// 提取目录路径
	dirPath := filepath.Dir(filePath)

	// 检查文件夹是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 文件夹不存在，创建文件夹
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
