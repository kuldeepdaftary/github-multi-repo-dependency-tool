package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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
			fmt.Println(depUrl)
			exists, val := checkDatabase(depUrl)
			if exists == true {
				storedPRData := &StoredPRData{}
				json.Unmarshal([]byte(val), &storedPRData)
				fmt.Print(storedPRData)

				urlTmp := strings.Replace(storedPRData.PrUrl, "https://github.com/", "", 1)
				repoName := strings.Split(urlTmp, "/")[1]

				go changePrStatus(storedPRData.PrUrl, "success", repoName, storedPRData.StatusUrl)
				go removeKey(githubDataPr.Pull_Request.Url)
			}
		} else {
			// DO NOTHING FOR NOW
		}
	}
}
