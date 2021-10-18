package parser

import (
	"regexp"
	"studyGolang/crawler/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

//解析信息
func ParseCity(contents []byte) engine.ParserResult {

	re := regexp.MustCompile(cityRe)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, c := range all {
		result.Items = append(result.Items, "User:"+string(c[2])) //用户名字
		name := string(c[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(c[1]),
			ParserFunc: func(bytes []byte) engine.ParserResult {
				return ParseProfile(bytes, name)
			},
		})
	}

	return result
}
