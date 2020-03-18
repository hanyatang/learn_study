package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

// 向city返回参数
type ParseResult struct {
	Request []Request
	Items   []interface{}
}

func Nilparser([]byte) ParseResult {
	return ParseResult{}
}
