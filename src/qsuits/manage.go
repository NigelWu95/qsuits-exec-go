package qsuits

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var modPath string

func Versions(path string) (versions []string, paths []string, err error) {

	err = filepath.Walk(filepath.Join(path, ".qsuits"), func(path string, f os.FileInfo, err error) error {
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

func WriteMod(path []string, version string) (isSuccess bool, err error) {

	if len(path) == 0 {
		return false, errors.New("no valid path")
	}
	modPath = filepath.Join(path[0], ".qsuits", "version.mod")
	modFile, err := os.Create(modPath)
	if err != nil {
		return false, err
	}

	var size int
	if len(path) == 1 {
		size, err = modFile.WriteString("version=" + version + ",path=" +
			filepath.Join(path[0], ".qsuits", "qsuits-" + version + ".jar"))
	} else {
		size, err = modFile.WriteString("version=" + version + ",path=" + path[1])
	}
	if err != nil {
		return false, err
	}
	if size <= 0 {
		return false, errors.New("no content wrote")
	}
	_ = modFile.Close()
	return true, nil
}

func ReadMod(path string) (version string, qsuitsPath string, err error) {

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
	modItems := strings.Split(string(bytes), ",")
	version = strings.Split(modItems[0], "=")[1]
	qsuitsPath = strings.Split(modItems[1], "=")[1]
	return version, qsuitsPath, nil
}

func LatestVersionFrom(versions []string) (latestVer string, latestVerNum int, err error) {

	if len(versions) == 0 {
		return latestVer, -1, errors.New("no versions")
	}
	var currentVer string
	vers := []string{"0", "0", "0"}
	var vNums []string
	var newV bool
	var vLen int
	for e := range versions {
		vNums = strings.Split(versions[e], ".")
		vLen = len(vNums)
		if vLen == 0 || vLen > 3 {
			return latestVer, -1, errors.New("error version: " + versions[e])
		} else if vLen == 1 {
			newV = false
			vers[0] = vNums[0]
			vers[1] = "0"
			vers[2] = "0"
		} else if vLen == 2 {
			newV = false
			vers[0] = vNums[0]
			seconds := vNums[1]
			if len(seconds) == 0 {
				vers[1] = "0"
				vers[2] = "0"
			} else if len(seconds) == 1 {
				vers[1] = seconds
				vers[2] = "0"
			} else {
				vers[1] = seconds[0:1]
				vers[2] = seconds[1:2]
			}
		} else {
			newV = true
			vers[0] = vNums[0]
			vers[1] = vNums[1]
			vers[2] = vNums[2]
		}
		if strings.EqualFold(vers[0], "") {
			vers[0] = "0"
		}
		if strings.EqualFold(vers[1], "") {
			vers[1] = "0"
		}
		if strings.EqualFold(vers[2], "") {
			vers[2] = "0"
		}
		if err != nil {
			return latestVer, -1, err
		}
		currentVer = strings.Join(vers, ".")
		if e == 0 || strings.Compare(currentVer, latestVer) > 0 {
			latestVer = currentVer
			latestVerNum = e
		} else if strings.Compare(currentVer, latestVer) == 0 {
			if newV || strings.Compare(strings.ReplaceAll(versions[e], ".", ""),
				strings.ReplaceAll(versions[latestVerNum], ".", "")) < 0 {
				latestVer = currentVer
				latestVerNum = e
			}
		}
		//fmt.Println(versions[e] + " -> " + currentVer)
	}
	return versions[latestVerNum], latestVerNum, nil
}

func Compare(version1 string, version2 string) (com int, err error) {

	latestVer, _, err := LatestVersionFrom([]string{version1, version2})
	if err != nil {
		return 0, err
	}
	if strings.Compare(latestVer, version1) == 0 {
		return 1, nil;
	} else {
		return -1, nil;
	}
}
