package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-boilerplate/pkg/server"
)

func init() {
	envFlag := flag.String("ENV", "development", "to define the environment")
	flag.Parse()

	err := os.Setenv("ENV", *envFlag)
	if err != nil {
		log.Fatal("Fail to set environment variable")
	}
}

func main() {
	c := InitConfig()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		os.Exit(0)
	}()

	server.ServeWeb(c.Server, getEnv())
}
