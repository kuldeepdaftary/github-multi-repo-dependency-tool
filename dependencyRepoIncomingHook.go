package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "sync"
)

func dependencyRepoIncomingHook(w http.ResponseWriter, r *http.Request) {
    fmt.Println("DEP")
    decoder := json.NewDecoder(r.Body)
	githubDataPr := &GithubDataPr{}
	err := decoder.Decode(&githubDataPr)

	if err != nil {
		panic(err)
	}

    if githubDataPr.Pull_Request.Merged == true {
        exists, url := checkDatabase(githubDataPr.Pull_Request.Url)
        if exists == true {
            var wg sync.WaitGroup

            changePrStatus(url, "success", "atlas-roku", &wg)
            removeKey(githubDataPr.Pull_Request.Url, &wg)

	        wg.Wait()
        }
    } else {
        // DO NOTHING FOR NOW
    }
}
