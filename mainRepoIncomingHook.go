package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "strings"
    "bytes"
)

var Statuses_Url string

func mainRepoIncomingHook(w http.ResponseWriter, r *http.Request) {

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

    prUrl := githubDataPr.Pull_Request.Url

    checkDependencyPr(pullUrl, prUrl)
}

func checkDependencyPr(pullUrl string, prUrl string) {
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

        state := githubDataDependency.State
        prRepoName := githubDataDependency.Head.Repo.Name

        if state == "closed" {
            status := "success"
            changePrStatus(pullUrl, status, prRepoName)
        } else {
            status := "failure"
            go updateDatabase(prUrl, pullUrl)
            go changePrStatus(pullUrl, status, prRepoName)
        }
    }
}

func changePrStatus(pullUrl string, status string, prRepoName string) {
    client := &http.Client{}

    message := prepareMessage(prRepoName, status)

    var jsonStr = []byte(`{"state": "` + status + `", "target_url": "` + pullUrl + `", "description": "` + message + `", "context": "Dependency Manager"}`)

    req, err := http.NewRequest("POST", Statuses_Url, bytes.NewBuffer(jsonStr))
    if err != nil {
		panic(err)
	}
    req.Header.Add("Authorization", `Basic dGNyYW5kczpiYWlsZXkxMjM=`)
    resp, err := client.Do(req)
    defer resp.Body.Close()

    if resp.StatusCode == 201 {
        fmt.Println("Pull Request Status Updated")
    } else {
        fmt.Println(resp.StatusCode)
    }
}

func prepareMessage(prRepoName string, status string) string {
    if status == "failure" {
        return "Requires: " + strings.ToUpper(prRepoName) + " (Click details ->)"
    } else if status == "success" {
        return "No Dependencies Found"
    } else {
        return "An Error has occured"
    }
}
