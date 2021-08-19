package main

import "fmt"

// App - the struct which contains pointers to database connections and etc....
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up our APP")
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
