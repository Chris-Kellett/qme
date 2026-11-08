package logger

import config "qme/Config"

func Debug(guildId string, logText string, a ...any) {
	if config.Settings.LOGGINGLEVEL <= config.LoggingLevelDebug {
		infoLine, formattedLogText := ParseLoggingText(guildId, logText, a...)
		SendLogs(infoLine, formattedLogText, config.LoggingLevelDebug, true)
	}
}
