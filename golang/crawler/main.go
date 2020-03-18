package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	// 获取网页的url
	engine.Run(engine.Request{
		Url:        "http://www.7799520.com/jiaou",
		ParserFunc: parser.ParseCityList,
	})

}
