package qsuits

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"qsuits-exec-go/src/utils"
	"strconv"
	"strings"
	"sync"
	"time"
)

type HttpGet struct {
	Url           string
	HttpClient    *http.Client
	//MediaType     string
	//MediaParams   map[string]string
	ContentLength int64
	DownloadBlock int64
	DownloadRange [][]int64
	Count         int
	TempFiles     []*os.File
	File          *os.File
	WG            sync.WaitGroup
}

type MavenSearchJson struct {
	//ResponseHeader struct {
	//	Status int `json:"status"`
	//	QTime int `json:"QTime"`
	//	Params struct {
	//		Q string `json:"q"`
	//		Core string `json:"core"`
	//		Indent string `json:"indent"`
	//		Spellcheck string `json:"spellcheck"`
	//		Fl string `json:"fl"`
	//		Start string `json:"start"`
	//		Sort string `json:"sort"`
	//		SpellcheckCount string `json:"spellcheck.count"`
	//		Rows string `json:"rows"`
	//		Wt string `json:"wt"`
	//		Version string `json:"version"`
	//	} `json:"params"`
	//} `json:"responseHeader"`
	Response struct {
		NumFound int `json:"numFound"`
		Start int `json:"start"`
		Docs []struct {
			ID string `json:"id"`
			G string `json:"g"`
			A string `json:"a"`
			LatestVersion string `json:"latestVersion"`
			RepositoryID string `json:"repositoryId"`
			P string `json:"p"`
			Timestamp int64 `json:"timestamp"`
			VersionCount int `json:"versionCount"`
			Text []string `json:"text"`
			Ec []string `json:"ec"`
		} `json:"docs"`
	} `json:"response"`
	//Spellcheck struct {
	//	Suggestions []interface{} `json:"suggestions"`
	//} `json:"spellcheck"`
}

func GetLatestVersion() (latestVersion string, err error) {
	latestVersion, err = GetLatestVersionByGithubProject()
	if err != nil {
		//fmt.Println(err.Error())
		latestVersion, err = GetLatestVersionBySearchMaven()
	}
	return latestVersion, err
}

func GetLatestVersionBySearchMaven() (latestVersion string, err error) {

	client := &http.Client{
		Timeout: time.Minute,
	}
	req, err := http.NewRequest("GET", "https://search.maven.org/solrsearch/select?q=a:qsuits&start=0&rows=20", nil)
	if err != nil {
		return latestVersion, err
	}
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return latestVersion, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close
	var f MavenSearchJson
	err = json.Unmarshal(body, &f)
	if err != nil {
		return latestVersion, err
	}
	return f.Response.Docs[0].LatestVersion, nil
}

func GetLatestVersionByGithubProject() (latestVersion string, err error) {

	client := &http.Client{
		Timeout: time.Minute,
	}
	req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/NigelWu95/qiniu-suits-java/master/pom.properties", nil)
	if err != nil {
		return latestVersion, err
	}
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return latestVersion, err
	}
	rd := bufio.NewReader(resp.Body)
	defer resp.Body.Close()
	for {
		line, err := rd.ReadString('\n') // 以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			return latestVersion, errors.New(fmt.Sprintf("get pom.properties version failed, %s", err.Error()))
		} else if strings.Contains(line, "version=") {
			return strings.Trim(strings.Split(line, "version=")[1], "\n"), nil
		}
	}
}

func ConcurrentDownload(url string, filepath string, blockSize int64, timeout time.Duration) (err error) {

	get := new(HttpGet)
	if timeout == 0 {
		get.HttpClient = new(http.Client)
	} else {
		get.HttpClient = &http.Client{
			Timeout: timeout,
		}
	}
	get.Url = url
	get.DownloadBlock = blockSize  // 1048576 = 1M

	req, err := http.NewRequest("GET", get.Url, nil)
	req.Header.Set("Range", "bytes=0-100")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	resp, err := get.HttpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 206 {
		respDump, _ := httputil.DumpResponse(resp, false)
		return errors.New(string(respDump))
	}
	//get.MediaType, get.MediaParams, _ = mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	contentRange := strings.Split(resp.Header.Get("Content-Range"), "/")
	if len(contentRange) < 2 {
		return errors.New("can not get content-range")
	}
	get.ContentLength, _ = strconv.ParseInt(contentRange[1], 10, 64)
	get.Count = int(math.Ceil(float64(get.ContentLength / get.DownloadBlock)))
	get.File, err = os.Create(filepath)
	if err != nil {
		return err
	}
	var rangeStart int64 = 0
	for i := 0; i < get.Count; i++ {
		if i != get.Count - 1 {
			get.DownloadRange = append(get.DownloadRange, []int64{rangeStart, rangeStart + get.DownloadBlock - 1})
		} else {
			// 最后一块
			get.DownloadRange = append(get.DownloadRange, []int64{rangeStart, get.ContentLength - 1})
		}
		rangeStart += get.DownloadBlock
		rangeI := fmt.Sprintf("%d-%d", get.DownloadRange[i][0], get.DownloadRange[i][1])
		tempFile, err := os.OpenFile(filepath + "." + rangeI, os.O_RDWR|os.O_APPEND, 0)
		if err != nil || tempFile == nil {
			tempFile, err = os.Create(filepath + "." + rangeI)
			if err != nil {
				if i > 0 {
					for j := 0; j < i; j++ {
						_ = get.TempFiles[j].Close()
					}
				}
				return err
			}
		} else {
			fi, err := tempFile.Stat()
			if err != nil {
				return err
			}
			if fi != nil {
				get.DownloadRange[i][0] += fi.Size()
			} else {
				err = errors.New(" no file info from: " + tempFile.Name())
				return err
			}
		}
		get.TempFiles = append(get.TempFiles, tempFile)
	//}
	//
	//for i, _ := range get.DownloadRange {
		get.WG.Add(1)
		go get.RangeDownload(filepath, i)
	}

	get.WG.Wait()
	if goroutineErr != nil {
		return goroutineErr
	}

	var copyTimes = 5
	for i := 0; i < get.Count; i++ {
		cnt, err := io.Copy(get.File, get.TempFiles[i])
		if err != nil || cnt < (get.DownloadRange[i][1] - get.DownloadRange[i][0] + 1) {
			if copyTimes > 0 && cnt <= 0 {
				i--
				copyTimes--
			} else if err == nil {
				for j := i; j < get.Count; j++ {
					_ = get.TempFiles[j].Close()
				}
				return errors.New(fmt.Sprintf("copy error size %d bytes", cnt))
			} else {
				for j := i; j < get.Count; j++ {
					_ = get.TempFiles[j].Close()
				}
				return err
			}
		} else {
			copyTimes = 5
			_ = get.TempFiles[i].Close()
		}
	}
	err = get.File.Close()
	if err == nil {
		for i := 0; i < get.Count; i++ {
			err := os.Remove(get.TempFiles[i].Name())
			if err != nil {
				log.Printf("Remove temp file %s error %v.\n", get.TempFiles[i].Name(), err)
			}
		}
	}
	return err
}

var goroutineErrStr string
var goroutineErr error
var lock = sync.Mutex{}
func ConcurrentDownloadWithRetry(url string, filepath string, blockSize int64, timeout time.Duration, retry int) (err error) {
	for i := 0; i < retry; i++ {
		goroutineErr = nil
		err = ConcurrentDownload(url, filepath, blockSize, timeout)
		if err != goroutineErr {
			return err
		}
		if goroutineErr == nil {
			return nil
		}
	}
	return goroutineErr
}

func (get *HttpGet) RangeDownload(filepath string, i int) {

	defer get.WG.Done()
	if get.DownloadRange[i][0] > get.DownloadRange[i][1] {
		return
	}

	defer func() {
		// 捕获协程中的 panic 信息
		if err := recover(); err != nil {
			lock.Lock()
			goroutineErrStr = fmt.Sprintf("range download failed %s", err)
			goroutineErr = errors.New(goroutineErrStr)
			lock.Unlock()
			//fmt.Println(err) // 输出 panic 信息
		}
	}()

	rangeI := fmt.Sprintf("%d-%d", get.DownloadRange[i][0], get.DownloadRange[i][1])
	req, err := http.NewRequest("GET", get.Url, nil)
	req.Header.Set("Range", "bytes=" + rangeI)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	resp, err := get.HttpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	} else {
		cnt, err := io.Copy(get.TempFiles[i], resp.Body)
		if err != nil {
			panic(err)
		}
		if cnt != int64(get.DownloadRange[i][1] - get.DownloadRange[i][0] + 1) {
			reqDump, _ := httputil.DumpRequest(req, false)
			respDump, _ := httputil.DumpResponse(resp, false)
			errStr := fmt.Sprintf("%d, expect %d-%d, but got %d.\nRequest: %s\nResponse: %s\n",
				resp.StatusCode, get.DownloadRange[i][0], get.DownloadRange[i][1], cnt, string(reqDump), string(respDump))
			err = errors.New(errStr)
			panic(err)
		}
	}
}

func StraightDownload(qsuitsFilePath string, url string, timeout time.Duration) (err error) {

	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		err = ioutil.WriteFile(qsuitsFilePath, body, 0755)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New(resp.Status)
	}
}

func Download(resultDir string, version string, isLatest bool) (qsuitsFilePath string, err error) {

	done := make(chan struct{})
	if isLatest {
		go utils.SixDotLoopProgress(done, "latest qsuits version: " + version + " is downloading")
	} else {
		go utils.SixDotLoopProgress(done, "qsuits version: " + version + " is downloading")
	}

	url := "https://github.com/NigelWu95/qiniu-suits-java/releases/download/v" + version + "/qsuits-" + version + ".jar"

	qsuitsDir := filepath.Join(resultDir, ".qsuits");
	//qsuitsDirInfo, err := os.Stat(qsuitsDir)
	//if os.IsNotExist(err) {
	//	err = os.MkdirAll(filepath.Join(resultDir, ".qsuits"), 0755)
	//	if err != nil {
	//		return "", err
	//	}
	//	qsuitsDirInfo, err = os.Stat(qsuitsDir)
	//}
	//if err != nil {
	//	return "", err
	//}
	//if !strings.HasPrefix(qsuitsDirInfo.Mode().String(), "drwx") {
	//	err = errors.New("qsuits path's mode: " + qsuitsDirInfo.Mode().String() + " is illegal")
	//}
	err = os.MkdirAll(qsuitsDir, 0755)
	if err != nil {
		return "", err
	}
	qsuitsFilePath = filepath.Join(qsuitsDir, "qsuits-" + version + ".jar")
	//err = StraightDownload(url, qsuitsFilePath)
	err = ConcurrentDownload(url, qsuitsFilePath, 1048576, 0)
	if err != nil && strings.Contains(err.Error(), "copy error size") {
		err = ConcurrentDownload(url, qsuitsFilePath, 1048576, 0)
	}
	if err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			err = errors.New("sorry, this old version is deprecated, not recommend you to use it")
		} else {
			fmt.Printf("\r%s", err.Error())
			fmt.Println("\rdownload is retrying from maven...")
			url = "https://search.maven.org/remotecontent?filepath=com/qiniu/qsuits/" +
				version + "/qsuits-" + version + "-jar-with-dependencies.jar"
			//err = StraightDownload(url, qsuitsFilePath)
			err = ConcurrentDownload(url, qsuitsFilePath, 1048576, 0)
			if err != nil && strings.Contains(err.Error(), "copy error size") {
				err = ConcurrentDownload(url, qsuitsFilePath, 1048576, 0)
			}
		}
	}
	done <- struct{}{}
	close(done)
	if err == nil {
		fmt.Println(" -> finished.")
	} else {
		fmt.Print("\r")
	}
	return qsuitsFilePath, err
}

func Update(path string, version string, isLatest bool) (qsuitsFilePath string, err error) {

	qsuitsJarPath := filepath.Join(path, ".qsuits", "qsuits-" + version + ".jar")
	fileInfo, err := os.Stat(qsuitsJarPath)
	if err == nil && !fileInfo.IsDir() {
		// it is already latest version
		//return qsuitsJarPath, errors.New("it is already latest version")
		return qsuitsJarPath, nil
	} else {
		return Download(path, version, isLatest)
	}
}
