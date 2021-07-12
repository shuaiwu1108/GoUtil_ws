package model

type DBInfo struct {
	Url    string `json:"url"`
	Port   int    `json:"port"`
	Name   string `json:"name"`
	Pass   string `json:"pass"`
	DbName string `json:"db_name"`
}
