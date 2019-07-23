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

	homePath, err := user.HomePath()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	qsuitsVersion, err := qsuits.GetLatestVersion()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	localVersion, localPath, err := qsuits.ReadMod(homePath)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println(err.Error())
		return
	}

	var qsuitsPath string
	if strings.Compare(qsuitsVersion, localVersion) > 0 {
		//qsuitsPath, updateErr := qsuits.Download(qsuitsVersion, homePath)
		qsuitsPath, err = qsuits.Update(qsuitsVersion, homePath)
		if err != nil {
			fmt.Println(err.Error())
			_, paths, err := qsuits.Versions(homePath)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("use current default version from local.")
			qsuitsPath = paths[0]
		}
	} else {
		qsuitsPath = localPath
	}
	if strings.Compare(localPath, "") == 0 {
		_, _ = qsuits.WriteMod(homePath, qsuitsVersion)
	}

	var params []string
	//flag.Parse()
	//params = flag.Args()
	params = os.Args[1:]
	if strings.Contains(qsuitsPath, "qsuits") {
		err = qsuits.Exec(qsuitsPath, strings.Join(params, " "))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		fmt.Printf("no valid qsuits path: %s\n", qsuitsPath)
	}
}
