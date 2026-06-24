<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/bot-telegram</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/bot-telegram"><img src="https://pkg.go.dev/badge/github.com/togo-framework/bot-telegram.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Telegram driver for the <a href="https://github.com/togo-framework/bot">togo bot</a> subsystem.</strong></p>
</div>

## Install

```bash
togo install togo-framework/bot
togo install togo-framework/bot-telegram
```

<!-- /togo-header -->

The **Telegram** driver for togo's [`bot`](https://github.com/togo-framework/bot)
subsystem, built on [go-telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api).
It long-polls Telegram and dispatches each message to the bot command/message
handlers you register once with `bot.OnCommand` / `bot.OnMessage`.

## Configure

Get a token from [@BotFather](https://t.me/BotFather), then set:

```bash
BOT_DRIVER=telegram
TELEGRAM_BOT_TOKEN=123456:ABC-DEF...
# TELEGRAM_POLL_TIMEOUT=30   # optional long-poll seconds
```

Blank-import the driver next to the base:

```go
import (
	_ "github.com/togo-framework/bot"
	_ "github.com/togo-framework/bot-telegram"
)
```

Register handlers (see the [bot](https://github.com/togo-framework/bot) README):

```go
bot.OnCommand("ping", func(ctx context.Context, b *bot.Service, m bot.Message) (string, error) {
	return "pong 🏓", nil
})
```

`m.Channel` is the numeric chat ID (as a string); `Service.Send(ctx, channel, msg)`
replies to it. Commands of the form `/cmd@YourBot` are matched as `cmd`.

## License

MIT © togo-framework

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
