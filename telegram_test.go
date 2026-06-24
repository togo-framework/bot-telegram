package telegram

import (
	"context"
	"os"
	"testing"

	"github.com/togo-framework/bot"
)

// The driver registers itself with the bot registry on import.
func TestDriverRegistered(t *testing.T) {
	found := false
	for _, n := range bot.Drivers() {
		if n == "telegram" {
			found = true
		}
	}
	if !found {
		t.Fatal("telegram driver not registered with bot")
	}
}

// Without a token the factory must error clearly (no panic, no nil bot).
func TestFactoryRequiresToken(t *testing.T) {
	old := os.Getenv("TELEGRAM_BOT_TOKEN")
	_ = os.Unsetenv("TELEGRAM_BOT_TOKEN")
	defer os.Setenv("TELEGRAM_BOT_TOKEN", old)

	b, err := makeDriver(nil, nil)
	if err == nil || b != nil {
		t.Fatalf("expected error without TELEGRAM_BOT_TOKEN, got (%v,%v)", b, err)
	}
}

func TestSendRejectsBadChannel(t *testing.T) {
	d := &driver{}
	if err := d.Send(context.Background(), "not-a-number", "hi"); err == nil {
		t.Fatal("expected error for non-numeric channel id")
	}
}
