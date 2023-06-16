package database

import (
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() {
	if err := mgm.SetDefaultConfig(nil, "links-shortener", options.Client().ApplyURI(os.Getenv("MONGOURI"))); err != nil {
		log.Fatal(err)
	}
}
