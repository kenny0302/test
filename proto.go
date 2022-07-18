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
	ModeB            bool  `json:"modeB"`
	VIP              []int `json:"vip"`
	ModeC            bool  `json:"modeC"`
	Rate             int   `json:"rate"`
	Point_Rate_Limit int   `json:"point_rate_limit"`
	ExtraMode        bool  `json:"extramode"`
	ExtraLimit       int   `json:"extralimit"`
	ExtraRate        int   `json:"extrarate"`
}
