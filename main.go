package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/HollischenBiber/TaskManager.git/app"
)

func main() {

	a := app.Application{}

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)

	err := a.Start()

	if err != nil {

		log.Fatal(err.Error())

	}

	<-ctx.Done()

	err = a.Stop()

	if err != nil {

		log.Fatal(err.Error())

	}
}
