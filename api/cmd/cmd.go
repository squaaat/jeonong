package cmd

import (
	"github.com/spf13/cobra"

	cmdgorm "github.com/squaaat/nearsfeed/api/cmd/gorm"
)

const (
	ArgEnv        = "environment"
	ArgEnvDefault = "alpha"
)

func newCliCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "j",
		Short: "nearsfeed-api application",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
}

func Start() error {
	c := newCliCmd()
	c.AddCommand(newHTTPCommand())
	c.AddCommand(cmdgorm.New())

	if err := c.Execute(); err != nil {
		return err
	}
	return nil
}
