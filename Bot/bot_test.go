package bot

import (
	config "qme/Config"
	"testing"
)

func Test_ValidBotToken(t *testing.T) {
	Init()
	if len(config.Settings.DISCORDSESSIONID) < 10 {
		t.Fatalf("Invalid Discord SessionID: [%s]", config.Settings.DISCORDSESSIONID)
	}
}
