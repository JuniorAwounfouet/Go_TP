package main

import (
	"minicrm/internal/app"
	"minicrm/internal/storage"
)

func main() {
	var store storage.Storer = storage.NewMemoryStore()
	app.Run(store)
}
