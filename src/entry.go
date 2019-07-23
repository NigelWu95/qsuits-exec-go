package main

import (
	"fmt"
	"os"
	"qsuits-exec-go/src/qsuits"
	"qsuits-exec-go/src/user"
	"strconv"
	"strings"
)

func main()  {

	homePath, err := user.HomePath()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, version, err := qsuits.CheckJavaRuntime()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("please install java first.")
		_, _, _ = qsuits.JdkDownload()
		return
	} else {
		vers := strings.Split(version, ".")
		var ver int
		if strings.EqualFold(vers[0], "1") {
			ver, err = strconv.Atoi(vers[1])
		} else {
			ver, err = strconv.Atoi(vers[0])
		}
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if ver < 8 {
			fmt.Println("please update your java to JDK1.8 or above")
			return
		}
	}

	var params []string
	params = os.Args[1:]
	if len(params) > 0 {
		op := params[0]
		if strings.EqualFold(op, "-Local") {
			qsuitsPath := localQsuitsPath(homePath)
			execQsuits(qsuitsPath, params[1:]);
		} else if strings.EqualFold(op, "clear") {
			clear(homePath)
		} else if strings.EqualFold(op, "chgver") {
			chgver(homePath, params)
		} else if strings.EqualFold(op, "help") ||
			strings.EqualFold(op, "--help") || strings.EqualFold(op, "-h") {
			help()
		} else {
			qsuitsPath := updatedQsuitsPath(homePath)
			execQsuits(qsuitsPath, params)
		}
	} else {
		help()
	}
}

func help() {
	fmt.Println("Usage:")
	fmt.Println("      this tool is a agent program for qsuits, your local environment " +
		"need java8 or above. In default mode, this tool will use latest java qsuits to exec, " +
		"you only need use qsuits-java's parameters to run. If you use local mode it mean you " +
		"dont want to update latest qsuits automatically.")
	fmt.Println("Options:")
	fmt.Println("        -Local          use current default qsuits version to exec.")
	fmt.Println("        --help/-h/help  print usage.")
	fmt.Println("Usage of qsuits:  https://github.com/NigelWu95/qiniu-suits-java")
}

func clear(homePath string) {
	versions, paths, err := qsuits.Versions(homePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	i := len(versions) - 1
	lastVersion := versions[i]
	result, err := qsuits.WriteMod(homePath, lastVersion)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !result {
		return
	}
	for i := range paths[0:i] {
		err := os.Remove(paths[i])
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func chgver(homePath string, params []string)  {
	if len(params) > 1 {
		ver := params[1]
		_, err := qsuits.Exists(homePath, ver)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result, err := qsuits.WriteMod(homePath, ver)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if !result {
			return
		}
	}
}

func localQsuitsPath(homePath string) string {

	_, qsuitsPath, err := qsuits.ReadMod(homePath)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	if os.IsNotExist(err) {
		versions, paths, err := qsuits.Versions(homePath)
		if err != nil {
			panic(err)
		}
		if len(versions) == 0 {
			fmt.Println("no qsuits in your local, so latest qsuits will be download...")
			qsuitsVersion, err := qsuits.GetLatestVersion()
			if err != nil {
				panic(err)
			}
			qsuitsPath, err = qsuits.Download(homePath, qsuitsVersion)
			if err != nil {
				panic(err)
			}
			result, err := qsuits.WriteMod(homePath, qsuitsVersion)
			if !result || err != nil {
				fmt.Println("write mode failed.")
				fmt.Println(err)
			} else {
				fmt.Println("set " + qsuitsVersion + " as default version.")
			}
		} else {
			qsuitsPath = paths[len(versions) - 1]
			fmt.Println("use local latest version.")
			return qsuitsPath
		}
	}
	return qsuitsPath
}

func updatedQsuitsPath(homePath string) string {

	qsuitsVersion, err := qsuits.GetLatestVersion()
	if err != nil {
		panic(err)
	}

	//qsuitsPath, updateErr := qsuits.Download(homePath, qsuitsVersion)
	qsuitsPath, err := qsuits.Update(homePath, qsuitsVersion)
	if err != nil {
		fmt.Println(err.Error())
		versions, paths, err := qsuits.Versions(homePath)
		if err != nil {
			panic(err)
		}
		if len(versions) == 0 {
			panic("no qsuits in local.")
		}
		qsuitsPath = paths[len(versions) - 1]
		fmt.Println("use local latest version.")
	}
	return qsuitsPath
}

func execQsuits(qsuitsPath string, params []string) {

	if strings.Contains(qsuitsPath, "qsuits") {
		err := qsuits.Exec(qsuitsPath, strings.Join(params, " "))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		fmt.Printf("no valid qsuits path: %s\n", qsuitsPath)
	}
}
