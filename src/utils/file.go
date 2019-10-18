package utils

import (
	"errors"
	"os"
)

func FileExists(path string) (isExists bool, err error) {

	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if fileInfo == nil {
		return false, errors.New("no file info")
	}
	if fileInfo.IsDir() {
		return true, errors.New(path + " is directory")
	} else {
		return true, nil
	}
}
