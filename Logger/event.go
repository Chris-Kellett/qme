package logger

import config "qme/Config"

func Event(guildId string, logText string, a ...any) {
	if config.Settings.LOGGINGLEVEL <= config.LoggingLevelEvent {
		infoLine, formattedLogText := ParseLoggingText(guildId, logText, a...)
		SendLogs(infoLine, formattedLogText, config.LoggingLevelEvent, true)
	}
}
