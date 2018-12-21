package main

import (
	"encoding/base64"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Environment variables options
var (
	TextEnvVarKeys       = [...]string{"t", "text"}
	MoveTypeEnvVarKeys   = [...]string{"m", "move"}
	HiddenTextEnvVarKeys = [...]string{"h", "hidden"}
)

// Application state configuration
type appConfig struct {
	text     string
	moveType string
	duration int
}

func main() {

	appConfig := readEnvVars()

	// Based on the move type,
	// perform the appropriate move
	switch appConfig.moveType {

	case "scroll":
		performScrolling(appConfig)

	case "blink":
		performBlinking(appConfig)

	default:
		panic("move is undefined")
	}

}

// Makes the text blink
func performBlinking(config appConfig) {
	figure.NewFigure(config.text, "", true).Blink(config.duration, 1000, 500)
	log.Fatal("Duration ended..")
}

// Makes the text scroll from right to left
func performScrolling(config appConfig) {
	postfixSpaces := strings.Builder{}
	space := " "
	for i := 0; i < len(config.text)*2; i++ {
		postfixSpaces.WriteString(space)
	}
	figure.NewFigure(fmt.Sprint(config.text, postfixSpaces.String()), "", true).Scroll(config.duration, 200, "right")
	log.Fatal("Duration ended..")
}

// Reads environment variables,
// It returns the application config based on the read environment variables
func readEnvVars() appConfig {
	// create default env config
	appConfig := appConfig{"Hala art", "scroll", int(math.MaxInt32)}

	// read the text
	for _, envKey := range TextEnvVarKeys {
		text, isSet := os.LookupEnv(envKey)
		if isSet {
			appConfig.text = text
		}
	}

	// read the move
	for _, envKey := range MoveTypeEnvVarKeys {
		moveType, isSet := os.LookupEnv(envKey)
		if isSet {
			appConfig.moveType = moveType
		}
	}

	// check if the text is hidden (base64 encoded)
	// if so, decode it
	for _, envKey := range HiddenTextEnvVarKeys {
		hiddenEnv, isSet := os.LookupEnv(envKey)
		if isSet {
			isHidden, err := strconv.ParseBool(hiddenEnv)
			if err != nil {
				break
			}
			decodedText, decodeErr := base64.StdEncoding.DecodeString(appConfig.text)
			if decodeErr == nil && isHidden {git
				appConfig.text = string(decodedText)
			}
		}
	}

	return appConfig
}
