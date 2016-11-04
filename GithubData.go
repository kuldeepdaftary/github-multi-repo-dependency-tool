package main

type GithubDataPr struct {
	Action       string `json:"action"`
	Pull_Request struct {
		Url   string `json:"url"`
		Title string `json:"title"`
		State string `json:"state"`
		User  struct {
			Login string `json:"login"`
		}
		Body         string `json:"body"`
		Statuses_Url string `json:"statuses_url"`
		Merged       bool   `json:"merged"`
	}
}

type GithubDataDependency struct {
	State string `json:"state"`
	Head  struct {
		Repo struct {
			Name string `json:"name"`
		}
	}
}

type StoredPRData struct {
	PrUrl     string `json:"prurl"`
	DepUrl    string `json:"depurl"`
	StatusUrl string `json:"statusurl"`
}
