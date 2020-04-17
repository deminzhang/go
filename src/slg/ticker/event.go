package Ticker

import (
	"common/event"
	"common/util"
	"log"
	"slg/const"
	"time"
)

func tickerSecond() {
	log.Println(">>ticking")
	for {
		time.Sleep(time.Second)
		Event.Call(Const.OnSecond, Util.MilliSecond())
	}
}
func tickerMinute() {
	log.Println(">>ticking")
	for {
		time.Sleep(time.Minute)
		Event.GoCall(Const.OnMinute, Util.MilliSecond())
	}
}

func init() {
	Event.Reg(Const.OnServerStart, func() {
		log.Println("Ticker.OnServerStart")
		go tickerSecond()
		go tickerMinute()
	})
	// Event.Reg(Const.OnSecond, func(mills int64) {
	// 	log.Println("Ticker.OnSecond")
	// })
	// Event.Reg(Const.OnMinute, func(mills int64) {
	// 	log.Println("Ticker.OnMinute")
	// })
}
