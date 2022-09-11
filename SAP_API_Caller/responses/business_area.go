package responses

type BusinessArea struct {
	D struct {
		Count   string `json:"__count"`
		Results []struct {
			BusinessArea string `json:"BusinessArea"`
			ToText       struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Text"`
		} `json:"results"`
	} `json:"d"`
}
