package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/phanimullapudi/go-rest-api-course/internal/transport/http"
)

// App - the struct which contains pointers to database connections and etc....
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up our APP")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go Rest API Couse")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
