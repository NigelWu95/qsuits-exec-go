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

	qsuitsUrl, err := qsuits.GetDownLoadUrl()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(qsuitsUrl)
	}

	fmt.Println(homePath)
	qsuitsPath, err := qsuits.Download(homePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(qsuitsPath)
	//resp, err := http.Get(qsuitsUrl)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//defer resp.Body.Close()
	//f, err := os.Create("qsuits-latest.jar")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//_, err = io.Copy(f, resp.Body)
}
