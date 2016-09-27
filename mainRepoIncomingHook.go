package main

import (
    "net/http"
    "encoding/json"
    "strings"
)

var Statuses_Url string

func mainRepoIncomingHook(w http.ResponseWriter, r *http.Request) {

    decoder := json.NewDecoder(r.Body)
	githubDataPr := &GithubDataPr{}
	err := decoder.Decode(&githubDataPr)

	if err != nil {
		panic(err)
	}

    if githubDataPr.Action == "opened" || githubDataPr.Action == "reopened"{
    	if strings.Contains(githubDataPr.Pull_Request.Body, "REQUIRED") {
            Statuses_Url = githubDataPr.Pull_Request.Statuses_Url
            depUrl, prUrl := githubDataPr.processRequiredBody()
            checkDependencyPr(depUrl, prUrl)
        }
    }
}

func (githubDataPr GithubDataPr) processRequiredBody() (string, string) {
    s := strings.Trim(strings.Split(githubDataPr.Pull_Request.Body, "REQUIRED:")[1], " ")

    newUrl := strings.Replace(s, "github.com", "api.github.com/repos", 1)
    depUrl := strings.Replace(newUrl, "pull", "pulls", 1)

    prUrlTmp := strings.Replace(githubDataPr.Pull_Request.Url, "api.github.com/repos", "github.com", 1)
    prUrl := strings.Replace(prUrlTmp, "pulls", "pull", 1)

    return depUrl, prUrl
}
