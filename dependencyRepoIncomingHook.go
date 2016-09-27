package main

import (
    // "fmt"
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

    if githubDataPr.Pull_Request.Merged == true {
	depUrl := strings.Replace(githubDataPr.Pull_Request.Url, "github.com", "api.github.com/repos", 1)
        exists, url := checkDatabase(depUrl)
        if exists == true {
            go changePrStatus(url, "success", "atlas-roku")
            go removeKey(githubDataPr.Pull_Request.Url)
        }
    } else {
        // DO NOTHING FOR NOW
    }
}
