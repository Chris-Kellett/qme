package main

import (
	"os"
	"os/signal"
	bot "qme/Bot"
	"syscall"
)

func TestingStart() {
	main()
}

func main() {
	bot.Init()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(shutdown)

	<-shutdown
}
