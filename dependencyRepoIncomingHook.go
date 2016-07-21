package main

import (
    "fmt"
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

    fmt.Println(githubDataPr.Pull_Request.State)
    fmt.Println(githubDataPr.Pull_Request.Merged)

}
