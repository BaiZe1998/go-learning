//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import "github.com/google/wire"

// wireApp init application.
func wireApp(url string) *App {
	wire.Build(NewMySQLClient, NewApp)
	return nil
}
