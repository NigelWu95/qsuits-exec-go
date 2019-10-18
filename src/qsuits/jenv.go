package qsuits

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"qsuits-exec-go/src/utils"
	"runtime"
	"strings"
)

func CheckJavaRuntime() (javaPath string, version string, err error) {
	javaPath, err = exec.LookPath("java")
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

func JdkDownload() (javaPath string, err error) {

	osName := runtime.GOOS
	osArch := runtime.GOARCH
	fmt.Printf("os: %s_%s\n", osName, osArch)
	jdkFileName := "jdk-8u231"

	if strings.Contains(osName, "darwin") {
		jdkFileName += "-macosx-x64.dmg"
	} else if strings.Contains(osName, "linux") {
		if strings.Contains(osArch, "64") {
			jdkFileName += "-linux-x64.tar.gz"
		} else if strings.Contains(osArch, "86") {
			jdkFileName += "-linux-i586.tar.gz"
		} else {
			err := errors.New("no tar.gz file to download of this go arch")
			return jdkFileName, err
		}
	} else if strings.Contains(osName, "windows") {
		if strings.Contains(osArch, "64") {
			jdkFileName += "-windows-x64.exe"
		} else if strings.Contains(osArch, "86") {
			jdkFileName += "-windows-i586.exe"
		} else {
			err := errors.New("no executable file to download of this go arch")
			return jdkFileName, err
		}
	} else {
		err := errors.New("no jdk to download of this go arch")
		return jdkFileName, err
	}

	done := make(chan struct{})
	go progress.SixDotLoop(done, "jdk-downloading")
	err = ConcurrentDownloadWithRetry("http://qsuits.nigel.net.cn/" + jdkFileName, jdkFileName, 5)
	if err != nil && strings.Contains(err.Error(), "copy error size") {
		err = ConcurrentDownloadWithRetry("http://qsuits.nigel.net.cn/" + jdkFileName, jdkFileName, 5)
	}
	done <- struct{}{}
	close(done)
	if err != nil {
		fmt.Println(" error from url: http://qsuits.nigel.net.cn/" + jdkFileName)
		return jdkFileName, err
	} else {
		fmt.Println(" -> succeed.")
		return jdkFileName, nil
	}
}
