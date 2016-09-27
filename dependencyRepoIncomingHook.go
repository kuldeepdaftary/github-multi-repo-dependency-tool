package main

import (
    "fmt"
    "net/http"
    "strings"
    "encoding/json"
)

func dependencyRepoIncomingHook(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
	githubDataPr := &GithubDataPr{}
	err := decoder.Decode(&githubDataPr)

	if err != nil {
		panic(err)
	}

	if githubDataPr.Action == "closed" {
        if githubDataPr.Pull_Request.Merged == true {

    	depUrlTmp := strings.Replace(githubDataPr.Pull_Request.Url, "api.github.com/repos", "github.com", 1)
    	depUrl := strings.Replace(depUrlTmp, "pulls", "pull", 1)

            exists, url := checkDatabase(depUrl)
            if exists == true {
                urlTmp := strings.Replace(url, "https://github.com/", "", 1)
                fmt.Println(urlTmp)
                repoName := strings.Split(urlTmp, "/")[1]
                fmt.Println(repoName)
                go changePrStatus(url, "success", repoName)
                go removeKey(githubDataPr.Pull_Request.Url)
            }
        } else {
            // DO NOTHING FOR NOW
        }
    }
}
