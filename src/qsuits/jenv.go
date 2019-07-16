package qsuits

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

func JdkDownload() (string, string, error) {
	fmt.Println("recommend one tool for you: https://github.com/linux-china/jenv")
	fmt.Println("you can use it to install java more easily, the steps like:")
	fmt.Println("1. curl -L -s get.jenv.mvnsearch.org | bash")
	fmt.Println("2. source $HOME/.jenv/bin/jenv-init.sh")
	fmt.Println("3. jenv ls java")
	fmt.Println("4. jenv install java <latest version>")
	fmt.Println("(please allow the tool to set latest version as default.)")
	var javaPath string
	var version string
	return javaPath, version, nil
}
