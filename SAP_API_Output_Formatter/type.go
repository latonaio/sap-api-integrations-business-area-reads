package sap_api_output_formatter

type BusinessAreaReads struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	BusinessArea  string `json:"businessarea"`
	Deleted       string `json:"deleted"`
}

type BusinessArea struct {
	BusinessArea string `json:"BusinessArea"`
	ToText       string `json:"to_Text"`
}

type Text struct {
	BusinessArea     string `json:"BusinessArea"`
	Language         string `json:"Language"`
	BusinessAreaName string `json:"BusinessAreaName"`
}
