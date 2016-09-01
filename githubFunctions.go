package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "strings"
    "bytes"
)

func checkDependencyPr(depUrl string, prUrl string) {
    client := &http.Client{}

    req, err := http.NewRequest("GET", depUrl, nil)
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
            changePrStatus(depUrl, status, prRepoName)
        } else {
            status := "failure"
            newUrl := strings.Replace(depUrl, "api.github.com/repos", "github.com", 1)
            newDepUrl := strings.Replace(newUrl, "pulls", "pull", 1)
            go updateDatabase(prUrl, newDepUrl)
            go changePrStatus(newDepUrl, status, prRepoName)
        }
    }
}

func changePrStatus(depUrl string, status string, prRepoName string) {
    client := &http.Client{}

    message := prepareMessage(prRepoName, status)

    var jsonStr = []byte(`{"state": "` + status + `", "target_url": "` + depUrl + `", "description": "` + message + `", "context": "Dependency Manager"}`)
    fmt.Println(bytes.NewBuffer(jsonStr))
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
