package main

import (
	"github.com/rs/zerolog/log"

	"github.com/squaaat/nearsfeed/api/cmd"
)

func main() {
	err := cmd.Start()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
