package global

// init functions for global.dev.go and global.prod.go global variables
type globalContextInitFunc func() error
