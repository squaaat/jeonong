package cmd

import (
	"net"

	"github.com/spf13/cobra"

	apphttp "github.com/squaaat/nearsfeed/api/internal/app/http"
	"github.com/squaaat/nearsfeed/api/internal/config"
	"github.com/squaaat/nearsfeed/api/internal/container"
)

func newHTTPCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "http",
		Short: "about nearsfeed-api http server",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}

	startC := &cobra.Command{
		Use:   "start",
		Short: "run http application",
		RunE: func(cmd *cobra.Command, _ []string) error {
			env, err := cmd.Flags().GetString(ArgEnv)
			if err != nil {
				return err
			}

			cfg, err := config.New(env)
			if err != nil {
				return err
			}

			cont, err := container.New(cfg)
			if err != nil {
				return err
			}
			host := net.JoinHostPort("0.0.0.0", cfg.ServerHTTP.Port)
			err = apphttp.New(cont).Listen(host)
			if err != nil {
				return err
			}
			return nil
		},
	}
	startC.Flags().String(ArgEnv, ArgEnvDefault, "set environment to run gorm cli")

	c.AddCommand(startC)

	return c
}
