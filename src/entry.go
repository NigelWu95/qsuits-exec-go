package main

import (
	"fmt"
	"qsuits-exec-go/src/qsuits"
	"qsuits-exec-go/src/user"
)

func main()  {

	javaPath, version, err := qsuits.CheckJavaRuntime()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("please install java first.")
		_, _, _ = qsuits.JdkDownload()
		return
	} else {
		fmt.Println(javaPath, version)
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
	} else {
		fmt.Println(qsuitsVersion)
	}

	fmt.Println(homePath)
	qsuitsPath, err := qsuits.Download(qsuitsVersion, homePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	qsuitsPath, err = qsuits.Update(qsuitsVersion, homePath)
	if err != nil {
		fmt.Println(err)
	}
	if qsuitsPath != "" {
		fmt.Println(qsuitsPath)
	}
}
