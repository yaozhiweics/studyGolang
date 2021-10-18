package main

import (
	"studyGolang/crawler/engine"
	"studyGolang/crawler/parser"
	"studyGolang/crawler/scheduler"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:url,
	//    ParserFunc: parser.ParseCityList,
	//    	})

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}
