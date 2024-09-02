//go:build !dev

package global

import "os"

var DevMode bool = false

// InitContext is called once at the start of the program to initialize global variables
var InitContext globalContextInitFunc = func() error {
	{
		devModeStr, ok := os.LookupEnv("DEV_MODE")
		if !ok {
			DevMode = false
		} else {
			DevMode = devModeStr == "true"
		}
	}
	return nil
}
