package main

type Config struct {
	UseProxy bool   `json:"use_proxy"`
	Proxy    string `json:"proxy"`
	ChatId   int64  `json:"chat_id"`
	Token    string `json:"token"`
}
