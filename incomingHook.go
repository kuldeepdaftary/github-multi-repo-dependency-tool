package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "strings"
)

var Statuses_Url string

func incomingHook(w http.ResponseWriter, r *http.Request) {

    decoder := json.NewDecoder(r.Body)
	githubDataPr := &GithubDataPr{}
	err := decoder.Decode(&githubDataPr)

	if err != nil {
		panic(err)
	}

	if strings.Contains(githubDataPr.Pull_Request.Body, "REQUIRED") {
        Statuses_Url = githubDataPr.Pull_Request.Statuses_Url
        githubDataPr.processRequiredBody()
    }
}

func (githubDataPr GithubDataPr) processRequiredBody() {
    s := strings.Trim(strings.Split(githubDataPr.Pull_Request.Body, "REQUIRED:")[1], " ")
    newUrl := strings.Replace(s, "github.com", "api.github.com/repos", 1)
    pullUrl := strings.Replace(newUrl, "pull", "pulls", 1)
    checkDependencyPr(pullUrl)
}

func checkDependencyPr(pullUrl string) {
    client := &http.Client{}

    req, err := http.NewRequest("GET", pullUrl, nil)
    if err != nil {
		panic(err)
	}
    req.Header.Add("Authorization", `Basic dGNyYW5kczpiYWlsZXkxMjM=`)
    resp, err := client.Do(req)
    defer resp.Body.Close()

    if resp.StatusCode == 200 {
        githubDataDependency := &GithubDataDependency{}
        decoder := json.NewDecoder(resp.Body)
        err := decoder.Decode(&githubDataDependency)
        if err != nil {
    		panic(err)
    	}
        if githubDataDependency.State == "open" {
            githubDataPr := GithubDataPr{}
            githubDataPr.lockMainPr(Statuses_Url)
        }
    }
}

func (githubDataPr GithubDataPr) lockMainPr(Statuses_Url string) {
    fmt.Println(Statuses_Url)
}
