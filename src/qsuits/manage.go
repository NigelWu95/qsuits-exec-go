package qsuits

import (
	"os"
	"path/filepath"
	"strings"
)

func Versions(path string) ([]string, error) {

	var versions []string
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
		}
		return nil
	})
	if err != nil {
		return versions, err
	} else {
		return versions, nil
	}
}
