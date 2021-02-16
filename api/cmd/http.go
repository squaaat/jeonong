package cmd

import (
	"net"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/squaaat/nearsfeed/api/internal/app"
	"github.com/squaaat/nearsfeed/api/internal/config"
	serverhttp "github.com/squaaat/nearsfeed/api/internal/server/http"
)

func newHTTPCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "http",
		Aliases: []string{"h"},
		Short:   "about nearsfeed-api http server",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
	c.AddCommand(newHTTPStartCommand())

	return c
}

func newHTTPStartCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "start",
		Short: "run http application",
	}
	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run http server")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}
		cfg := config.MustInit(env, false)
		app := app.New(cfg)
		http := serverhttp.New(app)
		host := net.JoinHostPort("0.0.0.0", cfg.ServerHTTP.Port)

		if err := http.Listen(host); err != nil {
			log.Fatal().Msgf("%v", err)
		}
	}

	return c
}
