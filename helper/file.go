package helper

import (
	"path"
	"strings"
)

func GetFileBaseName(fileName string) string {
	return path.Base(fileName)
}

func GetFileExtension(fileName string) string {
	return path.Ext(GetFileBaseName(fileName))
}

func GetFileBaseNameOnly(fileName string) string {
	return strings.TrimSuffix(GetFileBaseName(fileName), GetFileExtension(fileName))
}
