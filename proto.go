package main

type Member struct {
	Id     string `json:"id"`
	Vip    int    `json:"vip"`
	Points int    `json:"points"`
	Coins  int    `json:"coins"`
}

type Request struct {
	Id         string `json:"id"`
	Used_Point int    `json:"used_point"`
	Total      int    `json:"total"`
}

type Config struct {
	ModeA      bool  `json:"ModeA"`
	VIP        []int `json:"Vip"`
	ModeB      bool  `json:"ModeB"`
	Rate       int   `json:"rate"`
	ExtraMode  bool  `json:"ExtraMode"`
	ExtraLimit int   `json:"ExtraLimit"`
	ExtraRate  int   `json:"ExtraRate"`
}
