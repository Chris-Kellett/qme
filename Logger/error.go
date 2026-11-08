package logger

import config "hashbat-image-bot/Config"

func Error(guildId string, err error, a ...any) {
	if config.Settings.LOGGINGLEVEL <= config.LoggingLevelDebug {
		infoLine, formattedLogText := ParseLoggingText(guildId, err.Error(), a...)
		SendLogs(infoLine, formattedLogText, config.LoggingLevelError, true)
	}
}

func ErrorText(guildId string, message string, a ...any) {
	if config.Settings.LOGGINGLEVEL <= config.LoggingLevelDebug {
		infoLine, formattedLogText := ParseLoggingText(guildId, message, a...)
		SendLogs(infoLine, formattedLogText, config.LoggingLevelError, true)
	}
}
