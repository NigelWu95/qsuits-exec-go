package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func CheckJavaRuntime() (string, string, error) {
	var version string
	javaPath, err := exec.LookPath("java")
	if err != nil {
		return javaPath, version, err
	}
	cmd := exec.Command(javaPath, "-version")
	var versionRet string
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return javaPath, version, err
	} else {
		if out.Len() > 0 {
			versionRet = out.String()
		} else {
			versionRet = stderr.String()
		}
	}
	var versionFields []string
	versionFields = strings.Split(strings.Split(versionRet, "\n")[0], " ")
	version = strings.Trim(versionFields[2], "\"")
	return javaPath, version, nil
}

func JdkDownload() {}

func main()  {
	javaPath, version, err := CheckJavaRuntime()
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println(javaPath, version)
	}
}
