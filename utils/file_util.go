package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func CreateDir(filePath string) bool {
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func WriteFile(fileName string, content []byte) bool {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		log.Error(err)
		return false
	} else {
		write, err := f.Write(content)
		if err == nil && write > 0 {
			return true
		}
	}
	return false
}

func CheckFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return true
}

// 追加数据
func AppendFile(fileName string, content []byte) bool {
	f, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	_, err := f.Write(content)
	defer func() {
		_ = f.Close()
	}()
	return err == nil
}

func ReadFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Error(err)
		return nil
	}
	return file
}
