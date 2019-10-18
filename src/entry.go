package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/inconshreveable/go-update"
	"net/http"
	"os"
	"qsuits-exec-go/src/qsuits"
	"qsuits-exec-go/src/user"
	"qsuits-exec-go/src/utils"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main()  {

	homePath, err := user.HomePath()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, version, err := qsuits.CheckJavaRuntime()
	fmt.Println("do you want to download jdk8 now ? (yes/no)")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan();
	verify := scanner.Text()
	if strings.EqualFold("yes", verify) {
		jdkPath, err := qsuits.JdkDownload()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("jdk download as " + jdkPath)
		}
	}
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("please install java first, refer to https://blog.csdn.net/wubinghengajw/article/details/102612267.")
		fmt.Println("do you want to download jdk8 now ? (yes/no)")
		scanner := bufio.NewScanner(os.Stdin)
		verify := scanner.Text()
		if strings.EqualFold("yes", verify) {
			jdkPath, err := qsuits.JdkDownload()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("jdk download as " + jdkPath)
			}
		}
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
	length := len(params)
	if length > 0 {
		op1 := params[0]
		op2 := params[length - 1]
		if strings.EqualFold(op1, "--Local") || strings.EqualFold(op1, "-L") {
			qsuitsPath := localQsuitsPath(homePath)
			execQsuits(qsuitsPath, params[1:]);
		} else if strings.EqualFold(op2, "--Local") || strings.EqualFold(op2, "-L") {
			qsuitsPath := localQsuitsPath(homePath)
			execQsuits(qsuitsPath, params[0:length - 1]);
		} else if strings.EqualFold(op1, "selfupdate") || strings.EqualFold(op1, "upgrade") {
			selfUpdate()
		} else if strings.EqualFold(op1, "versions") {
			versions(homePath)
		} else if strings.EqualFold(op1, "clear") {
			clear(homePath)
		} else if strings.EqualFold(op1, "current") {
			currentVersion(homePath)
		} else if strings.EqualFold(op1, "chgver") {
			changeVersion(homePath, params)
		} else if strings.EqualFold(op1, "download") {
			download(homePath, params)
		} else if strings.EqualFold(op1, "help") ||
			strings.EqualFold(op1, "--help") || strings.EqualFold(op1, "-h") {
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
	fmt.Println("    this tool is a agent program for qsuits, your local environment " +
		"need java8 or above. In default mode, this tool will use latest java qsuits to exec, " +
		"you only need use qsuits-java's parameters to run. If you use local mode it mean you " +
		"dont want to update latest qsuits automatically.")
	fmt.Println("Options:")
	fmt.Println("        -Local/-L       use current default qsuits version to exec.")
	fmt.Println("        --help/-h/help  print usage.")
	fmt.Println("Commands:")
	fmt.Println("         help           print usage.")
	fmt.Println("         selfupdate     update this own executable program by itself.")
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
	latestVersion, latestNum, err := qsuits.LatestVersionFrom(versions)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result, err := qsuits.WriteMod([]string{homePath}, latestVersion)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if result {
		for i := range paths {
			if i == latestNum {
				continue
			}
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
	if err != nil {
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
		result, err := qsuits.WriteMod([]string{homePath}, ver)
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
			fmt.Printf("download %s failed.\n", ver)
			panic(err)
		} else {
			fmt.Printf("download %s succeeded.\n", ver)
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
				panic(err)
			}
			if len(versions) == 0 {
				fmt.Println("no qsuits in your local.")
				qsuitsVersion, err = qsuits.GetLatestVersion()
				if err != nil {
					fmt.Println(err.Error())
					return qsuitsPath
				}
				qsuitsPath, err = qsuits.Download(homePath, qsuitsVersion, true)
				if err != nil {
					panic(err)
				}
			} else {
				qsuitsVersion, num, err := qsuits.LatestVersionFrom(versions)
				if err != nil {
					panic(err)
				}
				qsuitsPath = paths[num]
				fmt.Println("use local latest version: " + qsuitsVersion)
			}
			result, err := qsuits.WriteMod([]string{homePath, qsuitsPath}, qsuitsVersion)
			if result && err == nil {
				fmt.Println("set " + qsuitsVersion + " as default local version.")
			} else {
				fmt.Println("write mode failed, " + err.Error())
			}
		} else {
			panic(err)
		}
	}
	return qsuitsPath
}

func updatedQsuitsPath(homePath string) string {

	qsuitsVersion, err := qsuits.GetLatestVersion()
	if err != nil {
		panic(err)
	}
	var versions []string
	var paths []string
	var versionsErr error
	var localLatestVer string
	var latestVerNum int
	versions, paths, versionsErr = qsuits.Versions(homePath)
	if versionsErr == nil {
		localLatestVer, latestVerNum, versionsErr = qsuits.LatestVersionFrom(versions)
		if versionsErr == nil {
			var com int
			com, versionsErr = qsuits.Compare(localLatestVer, qsuitsVersion)
			if com > 0 {
				fmt.Println("use local latest version: " + localLatestVer)
				return paths[latestVerNum]
			}
		}
	}
	qsuitsPath, err := qsuits.Update(homePath, qsuitsVersion, true)
	if err != nil {
		fmt.Println(err.Error() + ", update qsuits for version: " + qsuitsVersion + " failed.")
		if len(versions) == 0 {
			err = errors.New("no qsuits in your local")
			panic(err)
		}
		if versionsErr != nil {
			panic(versionsErr)
		}
		qsuitsPath = paths[latestVerNum]
		fmt.Println("use local latest version: " + localLatestVer)
	}
	return qsuitsPath
}

func execQsuits(qsuitsPath string, params []string) {

	if strings.Contains(qsuitsPath, "qsuits") {
		err := qsuits.Exec(qsuitsPath, params)
		if err != nil {
			//panic(err)
			fmt.Println(err.Error())
			return
		}
	} else {
		err := errors.New("invalid qsuits path: " + qsuitsPath)
		panic(err)
	}
}

func selfUpdate() {

	osName := runtime.GOOS
	osArch := runtime.GOARCH
	fmt.Printf("os: %s_%s\n", osName, osArch)
	binUrl := "https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_"

	if strings.Contains(osName, "darwin") {
		binUrl += "darwin_"
	} else if strings.Contains(osName, "linux") {
		binUrl += "linux_"
	} else if strings.Contains(osName, "windows") {
		binUrl += "windows_"
	} else {
		err := errors.New("no executable file to download of this go arch")
		panic(err)
	}

	if strings.Contains(osArch, "64") {
		binUrl += "amd64"
	} else if strings.Contains(osArch, "86") {
		binUrl += "386"
	} else {
		err := errors.New("no executable file to download of this go arch")
		panic(err)
	}

	if strings.Contains(osName, "windows") {
		binUrl += ".exe"
	}
	client := &http.Client{
		Timeout: 5 * time.Minute,
	}
	req, err := http.NewRequest("GET", binUrl, nil)
	if err != nil {
		panic(err)
	}
	done := make(chan struct{})
	go progress.SixDotLoop(done, "self-updating")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	done <- struct{}{}
	close(done)
	if err != nil {
		panic(err)
	}
	fmt.Println("-> succeed.")
}
