package responses

type ToText struct {
	D struct {
		Results []struct {
			BusinessArea     string `json:"BusinessArea"`
			Language         string `json:"Language"`
			BusinessAreaName string `json:"BusinessAreaName"`
		} `json:"results"`
	} `json:"d"`
}
