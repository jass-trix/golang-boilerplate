package main

import (
	"os"
)

const (
	development = "development"
	production  = "production"

	idevelopment = iota + 1
	iproduction
)

var (
	e  string
	ie int8
)

func init() {
	initEnv()
}

func initEnv() {
	e = os.Getenv("ENV")
	switch e {
	case production:
		ie = iproduction
	default:
		e = development
		ie = idevelopment
	}
}

// Get is used to get current environment from "TKPENV" variables on the OS
func getEnv() string {
	return e
}

// IsDevelopment is used to check whether current environment is on "development" environment or not
func isDevelopment() bool {
	return ie == idevelopment
}

// IsProduction is used to check whether current environment is on "production" environment or not
func isProduction() bool {
	return ie == iproduction
}
