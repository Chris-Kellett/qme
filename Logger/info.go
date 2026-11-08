package logger

import config "hashbat-image-bot/Config"

func Info(guildId string, logText string, a ...any) {
	if config.Settings.LOGGINGLEVEL <= config.LoggingLevelEvent {
		infoLine, formattedLogText := ParseLoggingText(guildId, logText, a...)
		SendLogs(infoLine, formattedLogText, config.LoggingLevelInfo, true)
	}
}
