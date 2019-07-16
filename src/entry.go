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

	path, err := user.Path()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(path)
	qsuitsUrl, err := qsuits.GetDownLoadUrl()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(qsuitsUrl)
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
