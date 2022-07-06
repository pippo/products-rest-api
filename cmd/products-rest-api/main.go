package main

import app "github.com/pippo/products-rest-api/internal/app/products-rest-api"

func main() {
	a := app.New()
	if err := a.Configure(); err != nil {
		panic(err.Error())
	}
	a.Run()
}
