package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

// Logging Variables ------------------------------------------------
const (
	LoggingLevelError = iota
	LoggingLevelWarn
	LoggingLevelEvent
	LoggingLevelInfo
	LoggingLevelDebug
)

type LoggingOptions struct {
	Name   string
	Colour Colour
}

var LoggingLevels map[int]LoggingOptions = map[int]LoggingOptions{
	LoggingLevelError: {
		Name:   "Error",
		Colour: Colours["red"],
	},
	LoggingLevelWarn: {
		Name:   "Warn",
		Colour: Colours["yellow"],
	},
	LoggingLevelEvent: {
		Name:   "Event",
		Colour: Colours["green"],
	},
	LoggingLevelInfo: {
		Name:   "Info",
		Colour: Colours["white"],
	},
	LoggingLevelDebug: {
		Name:   "Debug",
		Colour: Colours["blue"],
	},
}

// Colour Variables ------------------------------------------------
type Colour struct {
	Terminal         string `json:"Terminal,omitempty"`
	Html             string `json:"Html,omitempty"`
	GraphOpaque      string `json:"GraphOpaque,omitempty"`
	GraphTransparent string `json:"GraphTransparent,omitempty"`
}

var Colours map[string]Colour = map[string]Colour{
	"default": {
		Terminal:         "\033[0m",
		Html:             "#000000",
		GraphOpaque:      "rgba(0, 0, 0, 1)",
		GraphTransparent: "rgba(0, 0, 0, 0.2)",
	},
	"white": {
		Terminal:         "\033[97m",
		Html:             "#FFFFFF",
		GraphOpaque:      "rgba(, , , 1)",
		GraphTransparent: "rgba(, , , 0.2)",
	},
	"magenta": {
		Terminal:         "\033[35m",
		Html:             "#C30CC9",
		GraphOpaque:      "rgba(195, 12, 201, 1)",
		GraphTransparent: "rgba(195, 12, 201, 0.2)",
	},
	"yellow": {
		Terminal:         "\033[33m",
		Html:             "#FAF200",
		GraphOpaque:      "rgba(250, 242, 0, 1)",
		GraphTransparent: "rgba(250, 242, 0, 0.2)",
	},
	"green": {
		Terminal:         "\033[32m",
		Html:             "#28F200",
		GraphOpaque:      "rgba(40, 242, 0, 1)",
		GraphTransparent: "rgba(40, 242, 0, 0.2)",
	},
	"red": {
		Terminal:         "\033[31m",
		Html:             "#FF9EA0",
		GraphOpaque:      "rgba(242, 0, 8, 1)",
		GraphTransparent: "rgba(242, 0, 8, 0.2)",
	},
	"blue": {
		Terminal:         "\033[34m",
		Html:             "#25B7FF",
		GraphOpaque:      "rgba(0, 0, 255, 1)",
		GraphTransparent: "rgba(0, 0, 255, 0.2)",
	},
}

// App Constants & Variables ---------------------------------------
const (
	ROOT_FOLDER    string = "image-bot/"
	BOT_SUB_FOLDER string = "Bot/"
)

var (
	Session *discordgo.Session
)

// Environment Variables -------------------------------------------
var Settings Vars

type Vars struct {
	APPISDEV     bool
	APPROOTDIR   string
	APPSESSIONID string
	APPHOSTNAME  string

	DISCORDSESSIONID string
	DISCORDBOTTOKEN  string

	DB_NAME       string
	DB_USER       string
	DB_PASSWORD   string
	DB_IP_ADDRESS string
	DB_PORT       string

	LOGTODISCORD       bool
	LOGGINGCHANNELID   string
	LOGGINGUSESTHREADS bool
	LOGVERBOSESTACK    bool
	LOGFUNCTIONS       bool
	LOGGINGLEVEL       int
}

func init() {
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		fmt.Printf("wdErr ERROR: %s\n", wdErr.Error())
		panic(wdErr)
	}

	trimIndex := strings.LastIndex(wd, `image-bot\`)
	if trimIndex != -1 {
		Settings.APPROOTDIR = wd[:trimIndex+len(`image-bot\`)]
	} else {
		Settings.APPROOTDIR = wd
	}

	fmt.Printf("APPROOTDIR set as: [%s]\n", Settings.APPROOTDIR)
	loadErr := godotenv.Load(Settings.APPROOTDIR + ".env")
	if loadErr != nil {
		fmt.Printf("loadErr ERROR: %s\n", loadErr.Error())
		panic(loadErr)
	}

	envErr := parseEnvVariables()
	if envErr != nil {
		fmt.Printf("parseEnvVariables ERROR: %s\n", envErr.Error())
		panic(envErr)
	}

	currentHostName, err := os.Hostname()
	if err != nil {
		Settings.APPHOSTNAME = "Unknown"
	} else {
		Settings.APPHOSTNAME = currentHostName
	}

	Settings.APPSESSIONID = uuid.New().String()
}

func parseEnvVariables() error {
	var (
		err               error
		value             string
		notFoundText      string = "could not find .env value: "
		skippingText      string = "optional .env value not provided: "
		optionalErrorText string = "error processing optional .env value: "
	)

	value = os.Getenv("APPISDEV")
	if value == "" {
		return fmt.Errorf("%s%s", notFoundText, "APPISDEV")
	} else {
		Settings.APPISDEV, err = strconv.ParseBool(value)
		if err != nil {
			return err
		}
	}

	Settings.DISCORDBOTTOKEN = os.Getenv("DISCORDBOTTOKEN")
	if Settings.DISCORDBOTTOKEN == "" {
		return fmt.Errorf("%s%s", notFoundText, "DISCORDBOTTOKEN")
	}

	Settings.DB_NAME = os.Getenv("DB_NAME")
	if Settings.DB_NAME == "" {
		return fmt.Errorf("%s%s", notFoundText, "DB_NAME")
	}

	Settings.DB_USER = os.Getenv("DB_USER")
	if Settings.DB_USER == "" {
		return fmt.Errorf("%s%s", notFoundText, "DB_USER")
	}

	Settings.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	if Settings.DB_PASSWORD == "" {
		return fmt.Errorf("%s%s", notFoundText, "DB_PASSWORD")
	}

	Settings.DB_IP_ADDRESS = os.Getenv("DB_IP_ADDRESS")
	if Settings.DB_IP_ADDRESS == "" {
		return fmt.Errorf("%s%s", notFoundText, "DB_IP_ADDRESS")
	}

	Settings.DB_PORT = os.Getenv("DB_PORT")
	if Settings.DB_PORT == "" {
		return fmt.Errorf("%s%s", notFoundText, "DB_PORT")
	}

	value = os.Getenv("LOGTODISCORD")
	if value == "" {
		fmt.Printf("%s%s\n", skippingText, "LOGTODISCORD")
	} else {
		Settings.LOGTODISCORD, err = strconv.ParseBool(value)
		if err != nil {
			fmt.Printf("%s%s :: %s\n", optionalErrorText, "LOGTODISCORD", err)
		}
	}

	value = os.Getenv("LOGGINGCHANNELID")
	if value == "" {
		fmt.Printf("%s%s\n", skippingText, "LOGGINGCHANNELID")
	} else {
		Settings.LOGGINGCHANNELID = value
	}

	value = os.Getenv("LOGGINGUSESTHREADS")
	if value == "" {
		fmt.Printf("%s%s\n", skippingText, "LOGGINGUSESTHREADS")
	} else {
		Settings.LOGGINGUSESTHREADS, err = strconv.ParseBool(value)
		if err != nil {
			fmt.Printf("%s%s :: %s\n", optionalErrorText, "LOGGINGUSESTHREADS", err)
		}
	}

	value = os.Getenv("LOGVERBOSESTACK")
	if value == "" {
		fmt.Printf("%s%s\n", skippingText, "LOGVERBOSESTACK")
	} else {
		Settings.LOGVERBOSESTACK, err = strconv.ParseBool(value)
		if err != nil {
			fmt.Printf("%s%s :: %s\n", optionalErrorText, "LOGVERBOSESTACK", err)
		}
	}

	value = os.Getenv("LOGFUNCTIONS")
	if value == "" {
		fmt.Printf("%s%s\n", skippingText, "LOGFUNCTIONS")
	} else {
		Settings.LOGFUNCTIONS, err = strconv.ParseBool(value)
		if err != nil {
			fmt.Printf("%s%s :: %s\n", optionalErrorText, "LOGFUNCTIONS", err)
		}
	}

	value = os.Getenv("LOGGINGLEVEL")
	if value == "" {
		fmt.Printf("%s%s\n", skippingText, "LOGGINGLEVEL")
	} else {
		Settings.LOGGINGLEVEL, err = strconv.Atoi(value)
		if err != nil {
			fmt.Printf("%s%s :: %s\n", optionalErrorText, "LOGGINGLEVEL", err)
		}
	}

	return nil
}
