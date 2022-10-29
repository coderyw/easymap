package model

import "sync"

type TestStruct struct {
	A int     `json:"a"`
	B string  `json:"b"`
	C float64 `json:"c"`
	D bool    `json:"d"`
}

type Struct2 struct {
	DD string `json:"dd"`
}

type Resp360 struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		MatchId        string `json:"matchId"`
		MatchTimeStamp string `json:"matchTimeStamp"`
		MatchDate      string `json:"matchDate"`
		MatchTime      string `json:"matchTime"`
		CatId          string `json:"catId"`
		CatName        string `json:"catName"`
		LeagueId       string `json:"leagueId"`
		LeagueName     string `json:"leagueName"`
		HomeId         string `json:"homeId"`
		HomeName       string `json:"homeName"`
		HomeScore      string `json:"homeScore"`
		AwayId         string `json:"awayId"`
		AwayName       string `json:"awayName"`
		AwayScore      string `json:"awayScore"`
		LiveStreams    []struct {
			Id    string `json:"id"`
			Type  string `json:"type"`
			State bool   `json:"state"`
			Name  string `json:"name"`
		} `json:"liveStreams"`
		MatchState string `json:"matchState"`
	} `json:"data"`
	Else []interface{} `json:"else"`
}

type WaitGroupWrapper struct {
	sync.WaitGroup
}
