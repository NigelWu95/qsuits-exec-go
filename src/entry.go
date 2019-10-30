package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/inconshreveable/go-update"
	"io"
	"os"
	"path/filepath"
	"qsuits-exec-go/src/manual"
	"qsuits-exec-go/src/qsuits"
	"qsuits-exec-go/src/user"
	"qsuits-exec-go/src/utils"
	"runtime"
	"strings"
	"time"
)

var homePath string

func main()  {

	var err error
	homePath, err = user.HomePath()
	if err != nil {
		fmt.Printf("get home path failed, %s\n", err.Error())
		return
	}

	var op string
	var customJava bool
	var javaPath string
	var params []string
	var jvmParams []string
	local := true
	length := len(os.Args)
	if length > 1 {
		op = os.Args[1]
		if strings.EqualFold(op, "selfupdate") || strings.EqualFold(op, "upgrade") {
			selfUpdate()
		} else if strings.EqualFold(op, "versions") {
			versions()
		} else if strings.EqualFold(op, "clear") {
			clear()
		} else if strings.EqualFold(op, "current") {
			currentVersion()
		} else if strings.EqualFold(op, "chgver") {
			changeVersion(os.Args[1:])
		} else if strings.EqualFold(op, "download") {
			err = download(os.Args[1:])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		} else if strings.EqualFold(op, "update") {
			err = download(os.Args[1:])
			if err != nil {
				fmt.Println(err.Error())
				return
			} else {
				changeVersion(os.Args[1:])
			}
		} else if strings.EqualFold(op, "setjdk") {
			setJdk(os.Args[1:])
		} else if strings.EqualFold(op, "help") ||
			strings.EqualFold(op, "--help") || strings.EqualFold(op, "-h") {
			usage()
		} else if strings.EqualFold(op, "accounthelp") {
			manual.AccountUsage()
		} else if strings.EqualFold(op, "storagehelp") {
			manual.StorageUsage()
		} else if strings.EqualFold(op, "filehelp") {
			manual.FileUsage()
		} else if strings.EqualFold(op, "filterhelp") {
			manual.FilterUsage()
		} else if strings.EqualFold(op, "processhelp") {
			manual.ProcessUsage()
		} else {
			op = "exec"
			for i := 1; i < length; i++ {
				if strings.EqualFold(os.Args[i], "-u") {
					local = false
				} else if strings.EqualFold(os.Args[i], "--Local") || strings.EqualFold(os.Args[i], "-L") {
					local = true
				} else if strings.EqualFold(os.Args[i], "-j") || strings.EqualFold(os.Args[i], "--java") {
					customJava = true
					if (i + 1) < length && !strings.EqualFold(string(os.Args[i + 1][0]), "-") {
						javaPath, err = qsuits.SetJdkMod(homePath, os.Args[i + 1], 8)
						if err != nil {
							fmt.Println(err.Error())
							return
						}
						i++
					} else {
						javaPath, err = checkJava(true)
						if err != nil {
							fmt.Println(err.Error())
							fmt.Println("no jdk in your local setting.")
							javaInstall()
							return
						} else {
							fmt.Printf("your custom jdk path is %s.\n", javaPath)
						}
					}
				} else {
					if strings.HasPrefix(os.Args[i], "-X") {
						jvmParams = append(jvmParams, os.Args[i])
					} else {
						params = append(params, os.Args[i])
					}
				}
			}
		}
	} else {
		javaPath, err = checkJava(false)
		if err != nil {
			usage()
		} else {
			customJava = true
			op = "exec"
		}
	}

	if strings.EqualFold(op, "exec") {
		if !customJava {
			javaPath, err = checkJava(false)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("please install java 8 or above first.")
				javaInstall()
				return
			}
		}
		var qsuitsPath string
		if local {
			qsuitsPath, err = localQsuitsPath()
		} else {
			qsuitsPath, err = updatedQsuitsPath()
		}
		if err == nil {
			execQsuits(javaPath, qsuitsPath, jvmParams, params)
		} else {
			fmt.Println(err.Error())
		}
	}
}

func usage() {

	fmt.Println("Usage of qsuits:")
	fmt.Println("    this tool is a agent program for qsuits, your local environment " +
		"need java8 or above. In default mode, this tool will use latest java qsuits to exec, " +
		"you only need use qsuits-java's parameters to run. If you use local mode with \"-L/--Local\" it mean you " +
		"dont want to update latest qsuits automatically.")
	fmt.Println("Options:")
	fmt.Println("        -h                     Print usage.")
	fmt.Println("        -u                     Update latest qsuits version to exec.")
	fmt.Println("        -L/--Local             Use current default qsuits version to exec.")
	fmt.Println("        -j/--java [<jdkpath>]  Use custom jdk by existing setting or assigned <jdkpath>.")
	fmt.Println("Commands:")
	fmt.Println("         help                  Print usage.")
	fmt.Println("         selfupdate            Update this own executable program by itself.")
	fmt.Println("         versions              List all qsuits versions from local.")
	fmt.Println("         clear                 Remove all old qsuits versions from local.")
	fmt.Println("         current               Query local default qsuits version.")
	fmt.Println("         chgver <no.>          Set local default qsuits version.")
	fmt.Println("         download <no.>        Download qsuits with specified version.")
	fmt.Println("         update <no.>          Update qsuits with specified version.")
	fmt.Println("         setjdk <jdkpath>      Set jdk path as default.")
	fmt.Println()
	fmt.Println("Manual:")
	fmt.Println("         accounthelp           Print account usage")
	fmt.Println("         storagehelp           Print storage data source usage")
	fmt.Println("         filehelp              Print file data source usage")
	fmt.Println("         filterhelp            Print filter usage")
	fmt.Println("         processhelp           Print process usage")
	fmt.Println()
	fmt.Println("Usage of qsuits-java:          qsuits -path= -a= -process= -save-path= ...")
	fmt.Println("More details referer to:       https://github.com/NigelWu95/qiniu-suits-java")
}

func javaInstall() {
	fmt.Println("do you want to download jdk8 conformed to system now ? (yes/no)")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	verify := scanner.Text()
	if strings.EqualFold("y", verify) || strings.EqualFold("yes", verify) {
		jdkFileName, err := qsuits.JdkDownload()
		if err != nil {
			fmt.Printf("%s, maybe you need retry.\n", err.Error())
		} else {
			fmt.Printf("jdk download as %s, please install it refer to https://blog.csdn.net/wubinghengajw/article/details/102612267.\n", jdkFileName)
		}
	} else {
		fmt.Println("please install java refer to https://blog.csdn.net/wubinghengajw/article/details/102612267.")
	}
}

func versions() {

	vers, paths, err := qsuits.Versions(homePath)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for i := range vers {
			fmt.Printf("version: %s, path: %s\n", vers[i], paths[i])
		}
	}
}

func clear() {

	var latestVersion string
	var latestNum int
	var qsuitsPath string
	versions, paths, err := qsuits.Versions(homePath)
	if err == nil {
		latestVersion, latestNum, err = qsuits.LatestVersionFrom(versions)
	}
	if err == nil {
		qsuitsPath, err = qsuits.WriteMod([]string{homePath}, latestVersion)
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if qsuitsPath != "" {
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
		fmt.Printf("clear local versions succeeded, and set version: %s as local default.\n", latestVersion)
	} else {
		fmt.Printf("clear local old versions failed, but set version: %s as local default.\n", latestVersion)
	}
}

func currentVersion() {

	version, path, err := qsuits.ReadMod(homePath)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("version: %s, path: %s\n", version, path)
	}
}

func changeVersion(params []string) {

	if len(params) > 1 {
		ver := params[1]
		var qsuitsPath string
		_, err := utils.FileExists(filepath.Join(homePath, ".qsuits", "qsuits-" + ver + ".jar"))
		if err == nil {
			qsuitsPath, err = qsuits.WriteMod([]string{homePath}, ver)
		}
		if err != nil {
			fmt.Printf("chgver %s failed: %s \n", ver, err.Error())
			return
		}
		if qsuitsPath != "" {
			fmt.Printf("change local default version: %s succeeded.\n", ver)
		} else {
			fmt.Printf("change local default version: %s failed.\n", ver)
		}
	} else {
		fmt.Println("please chgver with version number like \"chgver 8.0.3\".")
	}
}

func download(params []string) (err error) {

	if len(params) > 1 {
		ver := params[1]
		_, err := qsuits.Download(homePath, ver, false)
		if err != nil {
			err = errors.New(fmt.Sprintf("download %s failed, %s", ver, err.Error()))
			return err
		} else {
			fmt.Printf("download %s succeeded.\n", ver)
			return nil
		}
	} else {
		err = errors.New("please set version number like \"download/update 8.0.3\"")
		return err
	}
}

func checkQsuitsVersionRecommend(version string) (err error) {

	url := "https://github.com/NigelWu95/qiniu-suits-java/releases/download/v" + version + "/qsuits-" + version + ".jar"
	err = qsuits.StraightHttpRequest(url, "HEAD", time.Minute, "")
	if err != nil && strings.Contains(err.Error(), "404 Not Found") {
		return errors.New(fmt.Sprintf("sorry, this old version: %s is deprecated, not recommend you to use it", version))
	} else {
		return nil
	}
}

func localQsuitsPath() (qsuitsPath string, err error) {

	var qsuitsVersion string
	qsuitsVersion, qsuitsPath, err = qsuits.ReadMod(homePath)
	if err != nil {
		versions, paths, err := qsuits.Versions(homePath)
		if !os.IsNotExist(err) {
			return qsuitsPath, errors.New(fmt.Sprintf("get local qsuits versions failed, %s", err.Error()))
		}
		if len(versions) == 0 {
			fmt.Println("no qsuits in your local.")
			qsuitsVersion, err = qsuits.GetLatestVersion()
			if err == nil {
				qsuitsPath, err = qsuits.Download(homePath, qsuitsVersion, true)
			}
			if err != nil {
				return qsuitsPath, errors.New(fmt.Sprintf("get latest qsuits failed, %s", err.Error()))
			}
		} else {
			var num int
			qsuitsVersion, num, err = qsuits.LatestVersionFrom(versions)
			if err != nil {
				return qsuitsPath, errors.New(fmt.Sprintf("get local qsuits versions failed, %s", err.Error()))
			}
			err = checkQsuitsVersionRecommend(qsuitsVersion)
			if err != nil {
				return paths[num], err
			}
			fmt.Printf("use local latest qsuits version: %s", qsuitsVersion)
		}
		qsuitsPath, err = qsuits.WriteMod([]string{homePath, qsuitsPath}, qsuitsVersion)
		if qsuitsPath != "" && err == nil {
			fmt.Printf(", and set %s as default local qsuits version.\n", qsuitsVersion)
		} else {
			fmt.Printf(", but set default local qsuits version failed, %s\n", err.Error())
		}
	} else {
		err = checkQsuitsVersionRecommend(qsuitsVersion)
		if err != nil {
			return qsuitsPath, err
		}
		fmt.Printf("use local qsuits version: %s\n", qsuitsVersion)
	}
	return qsuitsPath, nil
}

func updatedQsuitsPath() (qsuitsPath string, err error) {

	var setMode bool
	var versions []string
	var paths []string
	var versionsErr error
	var localLatestVer string
	var latestVerNum int
	qsuitsVersion, err := qsuits.GetLatestVersion()
	versions, paths, versionsErr = qsuits.Versions(homePath)
	localLatestVer, latestVerNum, versionsErr = qsuits.LatestVersionFrom(versions)
	if versionsErr == nil {
		if err != nil {
			fmt.Printf("get latest qsuits version failed, %s\n", err.Error())
			fmt.Printf("use local latest qsuits version: %s\n", localLatestVer)
			return paths[latestVerNum], nil
		} else {
			var com int
			com, versionsErr = qsuits.Compare(localLatestVer, qsuitsVersion)
			if com > 0 {
				fmt.Printf("use local latest qsuits version: %s\n", localLatestVer)
				return paths[latestVerNum], nil
			}
		}
	} else {
		setMode = true
		fmt.Printf("get local qsuits failed: %s\n", versionsErr.Error())
	}
	qsuitsPath, err = qsuits.Update(homePath, qsuitsVersion, true)
	if err != nil {
		output := fmt.Sprintf("update qsuits for version: %s failed, %s", qsuitsVersion, err.Error())
		if len(versions) == 0 {
			return qsuitsPath, errors.New(fmt.Sprintf("%s, but no qsuits in your local", output))
		}
		if versionsErr != nil {
			return qsuitsPath, errors.New(fmt.Sprintf("%s, but %s", output, versionsErr))
		}
		qsuitsPath = paths[latestVerNum]
		fmt.Printf("%s, use local latest qsuits version: %s.\n", output, localLatestVer)
	} else if setMode {
		isSuccess, err := qsuits.WriteMod([]string{homePath, qsuitsPath}, qsuitsVersion)
		if isSuccess != "" && err == nil {
			fmt.Printf("set %s as default local qsuits version.\n", qsuitsVersion)
		} else {
			fmt.Printf("set default local qsuits version failed, %s\n", err.Error())
		}
	}
	return qsuitsPath, nil
}

func execQsuits(javaPath string, qsuitsPath string, jvmParams []string, params []string) {

	if strings.Contains(qsuitsPath, "qsuits") {
		err := qsuits.Exec(javaPath, qsuitsPath, jvmParams, params)
		if err != nil {
			//fmt.Println("java path: ", javaPath)
			//fmt.Println("qsuits.jar path: ", qsuitsPath)
			fmt.Println(err.Error())
		}
	} else {
		fmt.Printf("invalid qsuits path: %s\n", qsuitsPath)
	}
}

func selfUpdate() {

	osName := runtime.GOOS
	osArch := runtime.GOARCH
	binUrl := "https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_"

	var isErr bool
	if strings.Contains(osName, "darwin") {
		if strings.Contains(osArch, "64") {
			binUrl += "darwin_amd64"
		} else {
			isErr = true
		}
	} else if strings.Contains(osName, "linux") {
		if strings.Contains(osArch, "64") {
			binUrl += "linux_amd64"
		} else if strings.Contains(osArch, "86") {
			binUrl += "linux_386"
		} else {
			isErr = true
		}
	} else if strings.Contains(osName, "windows") {
		if strings.Contains(osArch, "64") {
			binUrl += "windows_amd64.exe"
		} else if strings.Contains(osArch, "86") {
			binUrl += "windows_386.exe"
		} else {
			isErr = true
		}
	} else {
		isErr = true
	}
	if isErr {
		fmt.Printf("no executable file to download for this os_arch: %s_%s\n", osName, osArch)
		return
	}

	done := make(chan struct{})
	go utils.SixDotLoopProgress(done, "self-updating")

	qsuitsTempPath := filepath.Join(".", ".qsuitsselftemp")
	err := qsuits.ConcurrentDownloadWithRetry(binUrl, qsuitsTempPath, 1048576, 0, 2)
	if err == nil {
		var tempReader io.Reader
		tempReader, err = os.Open(qsuitsTempPath)
		if err == nil {
			err = update.Apply(tempReader, update.Options{})
		}
	}
	done <- struct{}{}
	close(done)
	if err != nil {
		fmt.Printf(", %s\n", err.Error())
		panic(err)
		return
	}
	_ = os.Remove(qsuitsTempPath)
	fmt.Println(" -> succeed.")
}

func setJdk(params []string) {

	if len(params) > 1 {
		jdkPath := params[1]
		isExists, err := utils.FileExists(jdkPath)
		if !isExists && err != nil {
			fmt.Printf("check jdk path failed, %s\n", err.Error())
		} else if isExists && err == nil {
			fmt.Println("jdk path must be a directory.")
		} else {
			_, err = qsuits.SetJdkMod(homePath, jdkPath, 8)
			if err != nil {
				fmt.Printf("set jdk path failed, %s\n", err.Error())
			} else {
				fmt.Println("set jdk path succeeded.")
			}
		}
	} else {
		fmt.Println("please setjdk with jdk path like \"jdk1.8.0_231\".")
	}
}

func checkJava(customJava bool) (javaPath string, err error) {

	javaPath = "java"
	var source = "system environment"
	var version string
	if customJava {
		javaPath, err = qsuits.GetJavaPathFromMod(homePath)
		if err == nil {
			source = "mod setting"
			_, err = utils.FileExists(javaPath)
			if err == nil {
				version, err = qsuits.GetJavaVersion(javaPath)
			}
		} else {
			err = errors.New(fmt.Sprintf("can not get custom java, %s.", err))
			return javaPath, err
		}
	} else {
		_, version, err = qsuits.CheckJavaRuntime()
		if err != nil {
			javaPath, err = qsuits.GetJavaPathFromMod(homePath)
			if err == nil {
				source = "mod setting"
				_, err = utils.FileExists(javaPath)
				if err == nil {
					version, err = qsuits.GetJavaVersion(javaPath)
				}
			} else {
				err = errors.New(fmt.Sprintf("no java in system and get custom java failed, %s.", err))
				return javaPath, err
			}
		}
	}
	if err == nil {
		err = qsuits.CheckJavaVersion(version, 8)
	}
	if err != nil {
		err = errors.New(fmt.Sprintf("get java path, %s from %s %s.", javaPath, source, err))
		return javaPath, err
	}
	return javaPath, nil
}
