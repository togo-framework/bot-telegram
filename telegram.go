// Package telegram is the Telegram driver for togo's bot subsystem. Blank-import
// it alongside github.com/togo-framework/bot and set BOT_DRIVER=telegram plus
// TELEGRAM_BOT_TOKEN to run a Telegram bot.
//
//	import _ "github.com/togo-framework/bot"
//	import _ "github.com/togo-framework/bot-telegram"
package telegram

import (
	"context"
	"fmt"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/togo-framework/bot"
	"github.com/togo-framework/togo"
)

func init() {
	bot.RegisterDriver("telegram", makeDriver)
}

func makeDriver(k *togo.Kernel, dispatch func(context.Context, bot.Message)) (bot.Bot, error) {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("bot-telegram: TELEGRAM_BOT_TOKEN is not set")
	}
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("bot-telegram: %w", err)
	}
	// Optional: longer poll timeout via TELEGRAM_POLL_TIMEOUT (seconds).
	timeout := 30
	if v := os.Getenv("TELEGRAM_POLL_TIMEOUT"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			timeout = n
		}
	}
	return &driver{api: api, dispatch: dispatch, timeout: timeout}, nil
}

type driver struct {
	api      *tgbotapi.BotAPI
	dispatch func(context.Context, bot.Message)
	timeout  int
	stop     chan struct{}
}

// Start long-polls Telegram for updates and dispatches each message.
func (d *driver) Start(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = d.timeout
	updates := d.api.GetUpdatesChan(u)
	d.stop = make(chan struct{})
	for {
		select {
		case <-ctx.Done():
			d.api.StopReceivingUpdates()
			return ctx.Err()
		case <-d.stop:
			d.api.StopReceivingUpdates()
			return nil
		case up, ok := <-updates:
			if !ok {
				return nil
			}
			msg := up.Message
			if msg == nil {
				continue
			}
			username := msg.From.UserName
			if username == "" {
				username = msg.From.FirstName
			}
			d.dispatch(ctx, bot.Message{
				Channel:  strconv.FormatInt(msg.Chat.ID, 10),
				User:     strconv.FormatInt(msg.From.ID, 10),
				Username: username,
				Text:     msg.Text,
				Platform: "telegram",
				Raw:      map[string]any{"update": up},
			})
		}
	}
}

// Stop ends the receive loop.
func (d *driver) Stop() error {
	if d.stop != nil {
		close(d.stop)
		d.stop = nil
	}
	return nil
}

// Send posts msg to a chat. channel is the numeric chat ID as a string.
func (d *driver) Send(ctx context.Context, channel, msg string) error {
	chatID, err := strconv.ParseInt(channel, 10, 64)
	if err != nil {
		return fmt.Errorf("bot-telegram: invalid chat id %q: %w", channel, err)
	}
	_, err = d.api.Send(tgbotapi.NewMessage(chatID, msg))
	return err
}
