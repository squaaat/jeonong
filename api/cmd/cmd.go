package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	cmdgorm "github.com/squaaat/jeonong/api/cmd/gorm"
)

const (
	ArgEnv        = "environment"
	ArgEnvShort   = "e"
	ArgEnvDefault = "alpha"
)

func newCliCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "j",
		Short: "jeonong-api application",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
}

func Start() {
	c := newCliCmd()
	c.AddCommand(newHTTPCommand())
	c.AddCommand(cmdgorm.New())

	if err := c.Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
