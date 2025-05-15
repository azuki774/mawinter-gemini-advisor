package fileoperator

import "os"

type FileOperator struct{}

func NewFileOperator() *FileOperator {
	return &FileOperator{}
}

func (f *FileOperator) LoadTxtFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	return string(data), err
}

func (f *FileOperator) WriteTxtFile(filePath string, data string) error {
	return os.WriteFile(filePath, []byte(data), 0644)
}
