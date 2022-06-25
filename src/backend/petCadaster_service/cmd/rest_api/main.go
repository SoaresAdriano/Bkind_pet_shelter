package main

import (
	"bkind_pet_shelter/src/backend/petCadaster_service/ports/rest"
	"context"
)

func main() {
	ctx := context.Background()
	server := rest.MustNewServer(rest.ServerParams{Port: 8080, App: app})

}
