package pangea

import "net/http"

type Application struct {
	server http.Server
}

func CreateApplication() *Application {
	return &Application{}
}

func (app *Application) Serve() {
	err := app.server.ListenAndServe()
	if err != nil {

	}
}
