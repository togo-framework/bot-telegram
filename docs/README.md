# bot-telegram — docs

**Telegram bot.** Telegram driver via telegram-bot-api — long-poll updates → the registry.

## Install

```bash
togo install togo-framework/bot-telegram
```

Registers on the [`bot`](https://github.com/togo-framework/bot) base; select it with **BOT_DRIVER (or bot.provider)**, then use **`togo bot`**.

## Interface

`Bot` — `Start`/`Stop`/`Send`, plus a command/handler registry (`OnCommand`/`OnMessage`) so any plugin can add bot commands.

## Configuration

| Env var | Description |
|---|---|
| `TELEGRAM_BOT_TOKEN` | Telegram bot token from @BotFather (required). |
| `TELEGRAM_POLL_TIMEOUT` |  |

## Usage & notes

Long-polls updates (timeout via `TELEGRAM_POLL_TIMEOUT`), routes commands to the registry, sends messages.

## Example

```bash
togo bot:send '#general' 'Deployed!'
togo bot:ask 'summarize the latest release'
```

## Links

- [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api)
- [Marketplace](https://to-go.dev/marketplace)
- [Source](https://github.com/togo-framework/bot-telegram)
