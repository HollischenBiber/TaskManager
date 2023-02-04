package app

import (
	"net/http"
	"os"

	serv "github.com/HollischenBiber/TaskManager.git/internal/adapters/http"
)

type Application struct {
	Server *http.Server
}

func (a *Application) Start() error {

	r := serv.InitRouter()

	port, _ := os.LookupEnv("TASKPORT")

	server := &http.Server{Addr: port, Handler: r}
	go server.ListenAndServe()

	a.Server = server

}

func (a *Application) Stop() error {

	return a.Server.Shutdown()

}
