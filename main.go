package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	http.HandleFunc("/pay", pay)
	http.ListenAndServe(":8000", nil)
}

func pay(w http.ResponseWriter, r *http.Request) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Can Not Read Body:" + err.Error()))
		return
	}

	var req Request
	err = json.Unmarshal(s, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Can Not Unmarshal Body:" + err.Error()))
		return
	}

	//取得會員資料
	val, err := rdb.Get(ctx, "member"+req.Id).Bytes()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No Member:" + err.Error()))
		return
	}

	var member Member
	err = json.Unmarshal(val, &member)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Can Not Unmarshal Data:" + err.Error()))
		return
	}

	//取得會員折扣 與 促銷模式
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Can Not Get Config Data:" + err.Error()))
		return
	}
	var final_total int
	final_total = req.Total
	//會員折扣
	if config.ModeA {
		switch member.Vip {
		case 1:
			final_total = int(math.Round(float64(req.Total*config.VIP[1])) / 100)
		case 2:
			final_total = int(math.Round(float64(req.Total*config.VIP[2])) / 100)
		case 3:
			final_total = int(math.Round(float64(req.Total*config.VIP[3])) / 100)
		default:
			final_total = req.Total
		}
	}

	//使用點數
	if config.ModeB {
		if req.Used_Point != 0 && member.Points >= req.Used_Point {
			final_total = final_total - (req.Used_Point * config.Rate)
		} else if req.Used_Point != 0 && member.Points < req.Used_Point {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No Enought Points"))
			return
		}
	}

	//額外活動
	if config.ExtraMode {
		if member.Vip != 0 && req.Used_Point > config.ExtraLimit {
			final_total = int(math.Round(float64(final_total*config.ExtraRate)) / 100)
		}
	}

	//執行扣款
	member.Points = member.Points - req.Used_Point
	member.Coins = member.Coins - final_total
	if member.Coins < 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No Enought Coins"))
		return
	}
	data, _ := json.Marshal(member)

	err = rdb.Set(ctx, "member"+member.Id, data, 0).Err()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Can Not Pay" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
