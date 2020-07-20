package qsuits

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
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
			ver := path[strings.Index(path, "qsuits-")+7:]
			if !strings.Contains(path, ".jar.") {
				ver = ver[0:strings.Index(ver, ".jar")]
			}
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

func WriteMod(path []string, version string) (qsuitsPath string, err error) {

	if len(path) == 0 {
		return qsuitsPath, errors.New("no valid path")
	}
	modPath = filepath.Join(path[0], ".qsuits", "version.mod")
	modFile, err := os.Create(modPath)
	if err != nil {
		return qsuitsPath, err
	}

	var size int
	if len(path) == 1 || path[1] == "" {
		qsuitsPath = filepath.Join(path[0], ".qsuits", "qsuits-"+version+".jar")
		size, err = modFile.WriteString("version=" + version + ",path=" + qsuitsPath)
	} else {
		qsuitsPath = path[1]
		size, err = modFile.WriteString("version=" + version + ",path=" + qsuitsPath)
	}
	if err != nil {
		return qsuitsPath, err
	}
	if size <= 0 {
		return qsuitsPath, errors.New("no content wrote")
	}
	_ = modFile.Close()
	return qsuitsPath, nil
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
	if len(modItems) != 2 {
		return version, qsuitsPath, errors.New("invalid version content in mode")
	}
	version = strings.Split(modItems[0], "=")[1]
	qsuitsPath = strings.Split(modItems[1], "=")[1]
	if version == "" || qsuitsPath == "" {
		return version, qsuitsPath, errors.New("invalid qsuits info in mode")
	}
	return version, qsuitsPath, nil
}

func LatestVersionFrom(versions []string) (latestVer string, latestVerNum int, err error) {

	if len(versions) == 0 {
		return latestVer, -1, errors.New("no versions")
	}
	var currentVer string
	vers := []int{0, 0, 0}
	lastest := []int{0, 0, 0}
	var vNums []string
	var newV bool
	var vLen int
	for e := range versions {
		vNums = strings.Split(versions[e], ".")
		vLen = len(vNums)
		if vLen == 0 {
			return latestVer, -1, errors.New("error version: " + versions[e])
		} else if vLen == 1 {
			newV = false
			vers[0], err = strconv.Atoi(vNums[0])
			vers[1] = 0
			vers[2] = 0
			currentVer = "0.0.0"
		} else if vLen == 2 {
			newV = false
			vers[0], err = strconv.Atoi(vNums[0])
			seconds := vNums[1]
			if len(seconds) == 0 {
				vers[1] = 0
				vers[2] = 0
				currentVer = vNums[0] + ".0.0"
			} else if len(seconds) == 1 {
				vers[1], err = strconv.Atoi(seconds)
				vers[2] = 0
				currentVer = vNums[0] + ".0." + seconds
			} else {
				vers[1], err = strconv.Atoi(seconds[0:1])
				vers[2], err = strconv.Atoi(seconds[1:2])
				currentVer = vNums[0] + "." + seconds[0:1] + "." + seconds[1:2]
			}
		} else {
			newV = true
			vers[0], err = strconv.Atoi(vNums[0])
			vers[1], err = strconv.Atoi(vNums[1])
			if strings.Contains(vNums[2], "-") { // -beta, -thin
				vers[2], err = strconv.Atoi(vNums[2][0:strings.Index(vNums[2], "-")])
				currentVer = vNums[0] + "." + vNums[1] + "." + vNums[2][0:strings.Index(vNums[2], "-")]
			} else {
				vers[2], err = strconv.Atoi(vNums[2])
				currentVer = vNums[0] + "." + vNums[1] + "." + vNums[2]
			}
		}
		if err != nil {
			return latestVer, -1, err
		}
		if !strings.Contains(versions[e], currentVer+".jar.") {
			if e == 0 || vers[0] > lastest[0] || (vers[0] == lastest[0] && vers[1] > lastest[1]) ||
				(vers[0] == lastest[0] && vers[1] == lastest[1] && vers[2] > lastest[2]) {
				lastest[0] = vers[0]
				lastest[1] = vers[1]
				lastest[2] = vers[2]
				latestVer = currentVer
				latestVerNum = e
			} else if strings.Compare(currentVer, latestVer) == 0 {
				if newV || strings.Compare(strings.ReplaceAll(versions[e], ".", ""),
					strings.ReplaceAll(versions[latestVerNum], ".", "")) < 0 {
					lastest[0] = vers[0]
					lastest[1] = vers[1]
					lastest[2] = vers[2]
					latestVer = currentVer
					latestVerNum = e
				}
			}
		}
		//fmt.Println(versions[e] + " -> " + currentVer)
	}
	if strings.EqualFold(latestVer, "") {
		return latestVer, latestVerNum, errors.New("no valid latest qsuits version")
	}
	return versions[latestVerNum], latestVerNum, nil
}

func Compare(version1 string, version2 string) (com int, err error) {

	latestVer, _, err := LatestVersionFrom([]string{version1, version2})
	if err != nil {
		return 0, err
	}
	if strings.Compare(latestVer, version1) == 0 {
		return 1, nil
	} else {
		return -1, nil
	}
}
