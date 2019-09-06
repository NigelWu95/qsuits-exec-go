package main

import (
	"errors"
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
		if strings.EqualFold(op, "-Local") || strings.EqualFold(op, "-L") {
			qsuitsPath := localQsuitsPath(homePath)
			execQsuits(qsuitsPath, params[1:]);
		} else if strings.EqualFold(op, "versions") {
			versions(homePath)
		} else if strings.EqualFold(op, "clear") {
			clear(homePath)
		} else if strings.EqualFold(op, "current") {
			currentVersion(homePath)
		} else if strings.EqualFold(op, "chgver") {
			changeVersion(homePath, params)
		} else if strings.EqualFold(op, "download") {
			download(homePath, params)
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
	fmt.Println("        -Local/-L       use current default qsuits version to exec.")
	fmt.Println("        --help/-h/help  print usage.")
	fmt.Println("Commands:")
	fmt.Println("         help           print usage.")
	fmt.Println("         versions       list all qsuits versions from local.")
	fmt.Println("         clear          remove all old qsuits versions from local.")
	fmt.Println("         current        query local default qsuits version.")
	fmt.Println("         chgver <no.>   set local default qsuits version.")
	fmt.Println("         download <no.> download qsuits with specified version.")
	fmt.Println("Usage of qsuits:  https://github.com/NigelWu95/qiniu-suits-java")
}

func versions(homePath string) {

	vers, paths, err := qsuits.Versions(homePath)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for i := range vers {
			fmt.Printf("version: %s, path: %s\n", vers[i], paths[i])
		}
	}
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
	if result {
		for i := range paths[0:i] {
			err := os.Remove(paths[i])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		fmt.Println("clear local versions succeeded.")
	} else {
		fmt.Println("clear local old versions failed.")
	}
}

func currentVersion(homePath string) {

	version, path, err := qsuits.ReadMod(homePath)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("version: %s, path: %s\n", version, path)
	}
}

func changeVersion(homePath string, params []string) {

	if len(params) > 1 {
		ver := params[1]
		_, err := qsuits.Exists(homePath, ver)
		if err != nil {
			fmt.Println("chgver " + ver + " failed: " + err.Error())
			return
		}
		result, err := qsuits.WriteMod(homePath, ver)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if result {
			fmt.Println("change local default version: " + ver + " succeeded.")
		} else {
			fmt.Println("change local default version: " + ver + " failed.")
		}
	} else {
		fmt.Println("please chgver with version number like \"chgver 7.1\".")
	}
}

func download(homePath string, params []string) {

	if len(params) > 1 {
		ver := params[1]
		_, err := qsuits.Download(homePath, ver, false)
		if err != nil {
			fmt.Println("download " + ver + " failed: " + err.Error())
		} else {
			fmt.Println("download " + ver + " succeeded.")
		}
	} else {
		fmt.Println("please download with version number like \"download 7.1\".")
	}
}

func localQsuitsPath(homePath string) string {

	_, qsuitsPath, err := qsuits.ReadMod(homePath)
	if err != nil {
		if os.IsNotExist(err) {
			var qsuitsVersion string
			versions, paths, err := qsuits.Versions(homePath)
			if err != nil {
				fmt.Println(err.Error())
				panic(err)
			}
			if len(versions) == 0 {
				err = errors.New("no qsuits in your local")
				fmt.Println(err.Error())
				panic(err)
			}
			if err != nil || len(versions) == 0 {
				if !os.IsNotExist(err) {
					fmt.Println(err.Error())
				} else {
					fmt.Println("no qsuits in your local.")
				}
				qsuitsVersion, err = qsuits.GetLatestVersion()
				if err != nil {
					fmt.Println(err.Error())
					return qsuitsPath
				}
				qsuitsPath, err = qsuits.Download(homePath, qsuitsVersion, true)
				if err != nil {
					fmt.Println(err.Error())
					return qsuitsPath
				}
			} else {
				i := len(versions) - 1
				qsuitsVersion = versions[i]
				qsuitsPath = paths[i]
				fmt.Println("use local latest version: " + qsuitsVersion)
			}
			i := len(versions) - 1
			qsuitsVersion = versions[i]
			qsuitsPath = paths[i]
			fmt.Println("use local latest version: " + qsuitsVersion)
			result, err := qsuits.WriteMod(homePath, qsuitsVersion)
			if !result || err != nil {
				fmt.Println("write mode failed, " + err.Error())
			} else {
				fmt.Println("set " + qsuitsVersion + " as default local version.")
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	return qsuitsPath
}

func updatedQsuitsPath(homePath string) string {

	qsuitsVersion, err := qsuits.GetLatestVersion()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	//qsuitsPath, updateErr := qsuits.Download(homePath, qsuitsVersion)
	qsuitsPath, err := qsuits.Update(homePath, qsuitsVersion, true)
	if err != nil {
		fmt.Println(err.Error() + ", update qsuits for version: " + qsuitsVersion + " failed.")
		versions, paths, err := qsuits.Versions(homePath)
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
		if len(versions) == 0 {
			err = errors.New("no qsuits in your local")
			fmt.Println(err.Error())
			panic(err)
		}
		qsuitsPath = paths[len(versions) - 1]
		fmt.Println("use local latest version: " + qsuitsVersion)
	}
	return qsuitsPath
}

func execQsuits(qsuitsPath string, params []string) {

	if strings.Contains(qsuitsPath, "qsuits") {
		err := qsuits.Exec(qsuitsPath, params)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		fmt.Printf("invalid qsuits path: %s\n", qsuitsPath)
	}
}
