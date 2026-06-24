---
name: bot-telegram
description: Run a Telegram bot in a togo app — configure BOT_DRIVER=telegram + TELEGRAM_BOT_TOKEN and register handlers with the bot plugin
---

# togo bot-telegram

Telegram driver for the togo `bot` subsystem.

## Setup

```bash
togo install togo-framework/bot
togo install togo-framework/bot-telegram
```

1. Create a bot with [@BotFather](https://t.me/BotFather), copy the token.
2. `.env`:
   ```bash
   BOT_DRIVER=telegram
   TELEGRAM_BOT_TOKEN=123456:ABC...
   ```
3. Register handlers with `bot.OnCommand` / `bot.OnMessage` (see the `bot` skill).

## Notes
- `m.Channel` is the numeric chat ID as a string; reply with `Service.Send`.
- `/cmd@YourBotName` is matched as `cmd`.
- The driver long-polls (no public webhook URL needed) — good for local dev and
  servers behind NAT. Tune with `TELEGRAM_POLL_TIMEOUT`.
- Never commit the token; keep it in `.env`.
