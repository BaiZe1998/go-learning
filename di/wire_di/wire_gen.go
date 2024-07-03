// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

// wireApp init application.
func wireApp(url string) *App {
	mySQLClient := NewMySQLClient(url)
	app := NewApp(mySQLClient)
	return app
}