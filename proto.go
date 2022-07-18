package main

type Member struct {
	Member_Id string `json:"member_id"`
	Vip       int    `json:"vip"`
	Points    int    `json:"points"`
	Coins     int    `json:"coins"`
}

type Request struct {
	Member_Id  string `json:"member_id"`
	Used_Point int    `json:"used_point"`
	Total      int    `json:"total"`
}

type Config struct {
	ModeB      bool  `json:"ModeB"`
	VIP        []int `json:"Vip"`
	ModeC      bool  `json:"ModeC"`
	Rate       int   `json:"rate"`
	ExtraMode  bool  `json:"ExtraMode"`
	ExtraLimit int   `json:"ExtraLimit"`
	ExtraRate  int   `json:"ExtraRate"`
}
