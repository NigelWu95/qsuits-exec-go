package qsuits

import (
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

	client := &http.Client{
		Timeout: time.Minute,
	}
	resp, err := client.Get("https://search.maven.org/solrsearch/select?q=a:qsuits&start=0&rows=20")
	if err != nil {
		return string(""), err
	}
	body, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close
	_ = resp.Close
	var f MavenSearchJson
	err = json.Unmarshal(body, &f)
	if err != nil {
		return string(""), err
	}
	return f.Response.Docs[0].LatestVersion, nil
}

func ConcurrentDownload(url string, filepath string) (err error) {

	get := new(HttpGet)
	get.HttpClient = new(http.Client)
	get.Url = url
	get.DownloadBlock = 1048576 // 1M

	req, err := http.NewRequest("GET", get.Url, nil)
	req.Header.Set("Range", "bytes=0-100")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	resp, err := get.HttpClient.Do(req)
	if err != nil {
		return err
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
	//}
	//
	//for i, _ := range get.DownloadRange {
		get.WG.Add(1)
		go get.RangeDownload(filepath, i)
	}

	get.WG.Wait()

	for i := 0; i < get.Count; i++ {
		cnt, err := io.Copy(get.File, get.TempFiles[i])
		if err != nil {
			for j := i; j < get.Count; j++ {
				_ = get.TempFiles[j].Close()
			}
			return err
		}
		if cnt != int64(get.DownloadRange[i][1] - get.DownloadRange[i][0] + 1) {
			for j := i; j < get.Count; j++ {
				_ = get.TempFiles[j].Close()
			}
			return errors.New("copy error size: " + string(cnt))
		}
		_ = get.TempFiles[i].Close()
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

func (get *HttpGet) RangeDownload(filepath string, i int) {

	if get.DownloadRange[i][0] > get.DownloadRange[i][1] {
		return
	}
	rangeI := fmt.Sprintf("%d-%d", get.DownloadRange[i][0], get.DownloadRange[i][1])

	defer func() {
		get.WG.Done()
		// 捕获协程中的 panic 信息
		if err := recover(); err != nil {
			errs++
			fmt.Println(err) // 输出 panic 信息
		}
	}()

	req, err := http.NewRequest("GET", get.Url, nil)
	req.Header.Set("Range", "bytes=" + rangeI)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	resp, err := get.HttpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	if err != nil {
		fmt.Printf("Download #%d failed.\n", i)
		panic(err)
	} else {
		rangeI := fmt.Sprintf("%d-%d", get.DownloadRange[i][0], get.DownloadRange[i][1])
		tempFile, err := os.OpenFile(filepath + "." + rangeI, os.O_RDWR|os.O_APPEND, 0)
		if err != nil || tempFile == nil {
			tempFile, _ = os.Create(filepath + "." + rangeI)
		} else {
			fi, _ := tempFile.Stat()
			if tempFile != nil {
				get.DownloadRange[i][0] += fi.Size()
			}
		}
		get.TempFiles = append(get.TempFiles, tempFile)
		cnt, err := io.Copy(tempFile, resp.Body)
		if err != nil {
			panic(err)
		}
		if cnt != int64(get.DownloadRange[i][1] - get.DownloadRange[i][0] + 1) {
			reqDump, _ := httputil.DumpRequest(req, false)
			respDump, _ := httputil.DumpResponse(resp, true)
			log.Panicf("Download error %d %v, expect %d-%d, but got %d.\nRequest: %s\nResponse: %s\n",
				resp.StatusCode, err, get.DownloadRange[i][0], get.DownloadRange[i][1], cnt, string(reqDump), string(respDump))
		}
	}
}

func StraightDownload(qsuitsFilePath string, url string) (err error) {

	client := &http.Client{
		Timeout: 10 * time.Minute,
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
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
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

var retry = 3
var errs = 0
func ConcurrentDownloadWithRetry()  {

}

func Download(resultDir string, version string, isLatest bool) (qsuitsFilePath string, err error) {

	done := make(chan struct{})
	if isLatest {
		go progress.SixDotLoop(done, "latest qsuits version: " + version + " is downloading")
	} else {
		go progress.SixDotLoop(done, "qsuits version: " + version + " is downloading")
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
	err = ConcurrentDownload(url, qsuitsFilePath)
	if err != nil {
		fmt.Println("\rdownload is retrying from maven...")
		url = "https://search.maven.org/remotecontent?filepath=com/qiniu/qsuits/" +
			version + "/qsuits-" + version + "-jar-with-dependencies.jar"
		//err = StraightDownload(url, qsuitsFilePath)
		err = ConcurrentDownload(url, qsuitsFilePath)
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

func Exists(path string, version string) (isExists bool, err error) {

	qsuitsJarPath := filepath.Join(path, ".qsuits", "qsuits-" + version + ".jar")
	fileInfo, err := os.Stat(qsuitsJarPath)
	if err != nil {
		return false, err
	}
	if fileInfo == nil {
		return false, errors.New("no file info")
	}
	if fileInfo.IsDir() {
		return true, errors.New(path + " is directory")
	} else {
		return true, nil
	}
}
