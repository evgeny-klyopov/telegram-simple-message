package telegram

import (
	"context"
	"golang.org/x/net/proxy"
	"gopkg.in/telegram-bot-api.v4"
	"net"
	"net/http"
	"strings"
)

func (c *Client) GetBot() *tgbotapi.BotAPI {
	return c.bot
}

func (c *Client) setBot() error {
	if c.config.UseProxy {
		setting := strings.Split(c.config.Proxy, "@")
		authData := strings.Split(setting[0], ":")

		dialer, err := proxy.SOCKS5(
			"tcp",
			setting[1],
			&proxy.Auth{User: authData[0], Password: authData[1]},
			proxy.Direct,
		)

		if err != nil {
			return err
		}

		client := &http.Client{Transport: &http.Transport{DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		}}}

		c.bot, err = tgbotapi.NewBotAPIWithClient(c.config.Token, client)
		if err != nil {
			return err
		}
	} else {
		var err error
		c.bot, err = tgbotapi.NewBotAPI(c.config.Token)
		if err != nil {
			return err
		}
	}

	return nil
}

