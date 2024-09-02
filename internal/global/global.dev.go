//go:build dev

package global

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const DevMode bool = true

// InitContext is called once at the start of the program to initialize global variables
var InitContext globalContextInitFunc = func() error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Non-JSON logging for dev
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	return nil
}
