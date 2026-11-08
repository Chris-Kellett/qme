package config

import (
	"testing"
)

func Test_ValidBotToken(t *testing.T) {
	if len(Settings.DISCORDBOTTOKEN) < 10 {
		t.Fatalf("Invalid BotToken: [%s]", Settings.DISCORDBOTTOKEN)
	}
}

func Test_ValidDbDetails(t *testing.T) {
	if len(Settings.DB_NAME) == 0 {
		t.Fatalf("Invalid DB_NAME: [%s]", Settings.DB_NAME)
	}
	if len(Settings.DB_USER) == 0 {
		t.Fatalf("Invalid DB_USER: [%s]", Settings.DB_USER)
	}
	if len(Settings.DB_PASSWORD) == 0 {
		t.Fatalf("Invalid DB_PASSWORD: [%s]", Settings.DB_PASSWORD)
	}
	if len(Settings.DB_IP_ADDRESS) == 0 {
		t.Fatalf("Invalid DB_IP_ADDRESS: [%s]", Settings.DB_IP_ADDRESS)
	}
	if len(Settings.DB_PORT) == 0 {
		t.Fatalf("Invalid DB_PORT: [%s]", Settings.DB_PORT)
	}
}

func Test_ValidHostName(t *testing.T) {
	if len(Settings.APPHOSTNAME) == 0 || Settings.APPHOSTNAME == "Unknown" {
		t.Fatalf("Invalid APPHOSTNAME: [%s]", Settings.APPHOSTNAME)
	}
}
