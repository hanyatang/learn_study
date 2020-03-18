package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<a class="name" href="(http://[^\s]+)" target="_blank">([^<]+)</a></h3>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User"+string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.Nilparser,
		})

	}
	return result
}
