package cmd

import (
	"log"

	"github.com/Brix101/nestfile/internal/settings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func generateKey() []byte {
	k, err := settings.GenerateKey()
	checkErr(err)
	return k
}
