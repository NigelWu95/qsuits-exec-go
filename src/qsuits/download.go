package qsuits

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

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

func GetLatestVersion() (string, error) {

	resp, err := http.Get("https://search.maven.org/solrsearch/select?q=a:qsuits&start=0&rows=20")
	if err != nil {
		return string(""), err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var f MavenSearchJson
	err = json.Unmarshal(body, &f)
	if err != nil {
		return string(""), err
	}
	return f.Response.Docs[0].LatestVersion, nil
}

func Download(version string, resultDir string) (string, error) {

	err := os.MkdirAll(filepath.Join(resultDir, ".qsuits"), os.ModePerm)
	if err != nil {
		return string(""), err
	}

	var jarFile string

	client := &http.Client{
		Timeout: 10 * time.Minute,
	}
	req, err := http.NewRequest("GET", "https://search.maven.org/remotecontent?filepath=com/qiniu/qsuits/" +
		version + "/qsuits-" + version + "-jar-with-dependencies.jar", nil)
	if err != nil {
		return jarFile, err
	}

	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")

	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return jarFile, err
	}
	err = resp.Body.Close()
	if err != nil {
		resp = nil
	}
	if resp.StatusCode == 200 {
		jarFile = filepath.Join(resultDir, ".qsuits/qsuits-" + version + ".jar")
		err = ioutil.WriteFile(jarFile, body, 0755)
		if err != nil {
			return jarFile, err
		}
		return jarFile, nil
	} else {
		return jarFile, errors.New(string(resp.StatusCode) + string(body))
	}
}

func Update(version string, resultDir string) (string, error) {

	qsuitsJarPath := filepath.Join(resultDir, ".qsuits/qsuits-" + version + ".jar")
	fileInfo, err := os.Stat(qsuitsJarPath)
	if err == nil && !fileInfo.IsDir() {
		return qsuitsJarPath, errors.New("it is already latest version")
	} else {
		return Download(version, resultDir)
	}
}
