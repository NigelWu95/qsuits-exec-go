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
		fmt.Println(err)
		return
	}

	qsuitsVersion, err := qsuits.GetLatestVersion()
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(homePath)
	//qsuitsPath, err := qsuits.Download(qsuitsVersion, homePath)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	qsuitsPath, err := qsuits.Update(qsuitsVersion, homePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	var params []string
	//flag.Parse()
	//params = flag.Args()
	params = os.Args[1:]
	fmt.Println(strings.Join(params, " "))
	if strings.Contains(qsuitsPath, "qsuits") {
		err = qsuits.Exec(qsuitsPath, strings.Join(params, " "))
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Printf("no valid qsuits path: %s\n", qsuitsPath)
	}
}
