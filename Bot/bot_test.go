package bot

import (
	config "hashbat-image-bot/Config"
	"testing"
)

func Test_ValidBotToken(t *testing.T) {
	Init()
	if len(config.Settings.DISCORDSESSIONID) < 10 {
		t.Fatalf("Invalid Discord SessionID: [%s]", config.Settings.DISCORDSESSIONID)
	}
}
