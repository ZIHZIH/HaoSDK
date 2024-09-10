package csv

import (
	"encoding/csv"
	"io"
	"os"
)

// CreateCsvFile 用于创建一个带有header的csv文件
func CreateCsvFile(fileName string, path string, header []string) error {
	// 判断目录
	err := os.MkdirAll(path, 0o777)
	if err != nil {
		return err
	}
	// 先删除在创建
	_ = os.Remove(path + fileName)

	f, err := os.Create(path + fileName) // 创建文件
	if err != nil {
		return err
	}
	_ = os.Chmod(path+fileName, os.ModePerm)
	defer func() {
		_ = f.Close()
	}()

	_, _ = f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	// 设置属性
	w.Comma = ','
	w.UseCRLF = true

	err = w.Write(header)
	if err != nil {
		return err
	}
	w.Flush()

	return nil
}

// WriteCsvFile 用于追加或者新建一个csv文件
func WriteCsvFile(fileName string, path string, data [][]string) error {
	if path != "" {
		// 判断目录
		err := os.MkdirAll(path, 0o777)
		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(path+fileName, os.O_RDWR|os.O_CREATE, 0o777)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	// 写入UTF-8 BOM
	_, err = f.WriteString("\xEF\xBB\xBF")
	if err != nil {
		return err
	}

	_, err = f.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	// 设置属性
	w.Comma = ','
	w.UseCRLF = true

	err = w.WriteAll(data)
	if err != nil {
		return err
	}

	w.Flush()

	return nil
}
