package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

func GetDownLoadUrl() (string, error) {

	var url string
	resp, err := http.Get("https://search.maven.org/solrsearch/select?q=a:qsuits&start=0&rows=20")
	body, err := ioutil.ReadAll(resp.Body)
	var f MavenSearchJson
	err = json.Unmarshal(body, &f)
	if err != nil {
		return url, err
	}
	return "https://search.maven.org/remotecontent?filepath=com/qiniu/qsuits/7.0-beta/qsuits-" +
		f.Response.Docs[0].LatestVersion + "-jar-with-dependencies.jar", nil
}

func main()  {

	qsuitsUrl, err := GetDownLoadUrl()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(qsuitsUrl)
	resp, err := http.Get(qsuitsUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	f, err := os.Create("qsuits-latest.jar")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = io.Copy(f, resp.Body)
}
