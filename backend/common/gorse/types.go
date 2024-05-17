package gorse

import "time"

type UserGoresBody struct {
	UserId string `json:"UserId"`
}

type PopularResp struct {
	Id    string  `json:"Id"`
	Score float64 `json:"Score"`
}

type VideosGoresBody struct {
	ItemId    string    `json:"ItemId"`
	Labels    []string  `json:"Labels"`
	Timestamp time.Time `json:"Timestamp"`
}
