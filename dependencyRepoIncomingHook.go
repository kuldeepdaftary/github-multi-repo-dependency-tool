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
	
	fmt.Println(githubDataPr)

    if githubDataPr.Pull_Request.Merged == true {
	fmt.Println("MERGED IS TRUE")
	depUrl := strings.Replace(githubDataPr.Pull_Request.Url, "api.github.com/repos", "github.com", 1)
	fmt.Println(depUrl) 
        exists, url := checkDatabase(depUrl)
        if exists == true {
	    fmt.Println("EXISTS IS TRUE")
	    fmt.Println(url)
            go changePrStatus(url, "success", "tv")
            go removeKey(githubDataPr.Pull_Request.Url)
        }
    } else {
        // DO NOTHING FOR NOW
    }
}
