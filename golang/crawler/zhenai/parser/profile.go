package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

// 预先编译
var ageRe = regexp.MustCompile(`<a  href="https://tao.yuemei.com/[0-9]/" class="part-link  ">([^<]+)<i></i></a>`)
var marriageRe = regexp.MustCompile(`<a  href="https://tao.yuemei.com/[0-9]/" class="part-link  ">([^<]+)<i></i></a>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	// 数字的提取
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	// 字符串的提取
	profile.Marriage = extractString(contents, marriageRe)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
