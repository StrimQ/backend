package logging

import (
	"github.com/rs/zerolog"
)

// ConfigureLogging creates a Zerolog instance based on app mode
func ConfigureLogging(debug bool) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
