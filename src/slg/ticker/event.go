package Ticker

import (
	"common/event"
	"common/utilX"
	"log"
	"slg/const"
	"time"
)

func ticker() {
	log.Println(">>ticker")
	for {
		time.Sleep(time.Millisecond * 300)
		Event.Call(Const.OnTick, utilX.MilliSecond())
	}
}

func tickerSecond() {
	log.Println(">>tickerSecond")
	for {
		time.Sleep(time.Second)
		Event.Call(Const.OnSecond, utilX.MilliSecond())
	}
}

func tickerMinute() {
	log.Println(">>tickerMinute")
	for {
		time.Sleep(time.Minute)
		Event.GoCall(Const.OnMinute, utilX.MilliSecond())
	}
}

func init() {
	Event.Reg(Const.OnServerStart, func() {
		log.Println("Ticker.OnServerStart")
		go ticker()
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
