package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
)

func GetTargetFilePathList(objectPath string, suffixArray []string) []string {
	var fileList []string
	files, _ := ioutil.ReadDir(objectPath)
	for _, file := range files {
		for _, suffix := range suffixArray {
			if suffix == path.Ext(file.Name()) {
				fileList = append(fileList, objectPath+file.Name())
			}
		}
	}
	return fileList
}

func ReadFileToString(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("read file fail:", err)
	}
	return string(data)
}

func WriteStringToFile(filePath string, data string) {
	err := ioutil.WriteFile(filePath, []byte(data), 'W')
	if err != nil {
		log.Fatal("rite data into file failed: ", err)
	}
}

func main() {
	filePath := "./"
	fileList := GetTargetFilePathList(filePath, []string{".txt", ".html"})
	for _, value := range fileList {
		fmt.Println(ReadFileToString(value))
	}
	WriteStringToFile(filePath+"hello.txt", "hello")
}
