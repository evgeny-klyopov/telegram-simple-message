package main

import (
	"gopkg.in/telegram-bot-api.v4"
)

type Client struct {
	config             Config
	bot                *tgbotapi.BotAPI
	defaultTypeMessage string
}

func NewClient(cfg Config) (*Client, error) {
	var client Client
	client = Client{
		config:             cfg,
		defaultTypeMessage: MarkdownTypeMessage,
	}
	err := client.setBot()
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *Client) Send(message string, typeMessage string) error {
	msg := tgbotapi.NewMessage(c.config.ChatId, message)
	msg.ParseMode = c.getTypeMessage(typeMessage)
	_, err := c.bot.Send(msg)

	if err != nil {
		return err
	}

	return nil
}