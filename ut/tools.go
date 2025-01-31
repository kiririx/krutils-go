package ut

import (
	"github.com/kiririx/krutils/tools"
)

func Convert(v any) *tools.Value {
	return tools.Convert(v)
}

func Algorithm() *tools.Algorithm {
	return &tools.Algorithm{}
}

func Async() *tools.Async {
	return &tools.Async{}
}

func Time() *tools.Time {
	return &tools.Time{}
}

func String() *tools.String {
	return &tools.String{}
}

func File() *tools.File {
	return &tools.File{}
}

func Regex() *tools.Regex {
	return &tools.Regex{}
}

func JSON() *tools.JSON {
	return &tools.JSON{}
}

func Struct() *tools.Struct {
	return &tools.Struct{}
}

func HttpClient() *tools.HttpClient {
	return tools.NewHttpClient()
}

func OpenAI(conf tools.OpenAI) *tools.OpenAI {
	return tools.NewOpenAI(conf)
}
