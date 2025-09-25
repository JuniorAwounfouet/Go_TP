package main

import (
	"github.com/JuniorAwounfouet/go-crm-app/internal/app"
	"github.com/JuniorAwounfouet/go-crm-app/internal/storage"
)

func main() {

	var store storage.Storer = storage.NewMemoryStore()
	app.Run(store)

}
