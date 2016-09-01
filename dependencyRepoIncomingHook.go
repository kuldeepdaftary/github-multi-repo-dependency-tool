package main

import (
    // "fmt"
    "net/http"
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
        exists, url := checkDatabase(githubDataPr.Pull_Request.Url)
        if exists == true {
            go changePrStatus(url, "success", "atlas-roku")
            go removeKey(githubDataPr.Pull_Request.Url)
        }
    } else {
        // DO NOTHING FOR NOW
    }
}
