package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.7799520.com/jiaou/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 拿到engine两个结构体的信息
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "city"+string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})

	}
	return result
}
