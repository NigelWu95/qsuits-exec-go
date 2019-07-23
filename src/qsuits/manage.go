package qsuits

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var modPath string

func Versions(path string) ([]string, []string, error) {

	var versions []string
	var paths []string
	err := filepath.Walk(filepath.Join(path, ".qsuits"), func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if strings.Contains(path, "qsuits-") {
			ver := strings.Trim(strings.Split(path, "qsuits-")[1], ".jar")
			versions = append(versions, ver)
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return versions, paths, err
	} else {
		return versions, paths, nil
	}
}

func WriteMod(path string, version string) (bool, error) {

	modPath = filepath.Join(path, ".qsuits", "version.mod")
	modFile, err := os.Create(modPath)
	if os.IsExist(err) {
		modFile, err = os.Open(modPath)
	}
	if err != nil {
		return false, err
	}
	size, err := modFile.WriteString("version=" + version + ",path=" +
		filepath.Join(path, ".qsuits", "qsuits-"+version+".jar"))
	if err != nil {
		return false, err
	}
	if size <= 0 {
		return false, errors.New("no content wrote")
	}
	_ = modFile.Close()
	return true, nil
}

func ReadMod(path string) (string, string, error) {

	modPath = filepath.Join(path, ".qsuits", "version.mod")
	modFile, err := os.Open(modPath)
	if err != nil {
		return "", "", err
	}
	bytes, err := ioutil.ReadAll(modFile)
	if err != nil {
		return "", "", err
	}
	_ = modFile.Close()
	modIterms := strings.Split(string(bytes), ",")
	version := strings.Split(modIterms[0], "=")[1]
	qsuitsPath := strings.Split(modIterms[1], "=")[1]
	return version, qsuitsPath, nil
}
