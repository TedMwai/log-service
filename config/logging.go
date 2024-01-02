package config

import "github.com/rs/zerolog"

func initLogger() {

	if AppIsDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		// UNIX Time is faster and smaller than most timestamps
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}