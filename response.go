package main

type response struct {
	Code rune              `json:"code"`
	Msg  string            `json:"msg"`
	Data map[string]string `json:"data"`
}
