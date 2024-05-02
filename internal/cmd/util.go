package cmd

import (
	"log"

	"github.com/Brix101/nestfile/internal/settings"
)


func generateKey() []byte {
	k, err := settings.GenerateKey()
	checkErr(err)
	return k
}
