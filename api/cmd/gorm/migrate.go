package gorm

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/squaaat/jeonong/api/internal/app"
	"github.com/squaaat/jeonong/api/internal/config"
	"github.com/squaaat/jeonong/api/migrations"
)

func newGormMigrate() *cobra.Command {
	c := &cobra.Command{
		Use:     "migrate",
		Aliases: []string{"m"},
		Short:   "it's about gorm migrate",
	}
	c.Run = func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	}

	c.AddCommand(newGormMigrateSync())
	c.AddCommand(newGormMigrateCreate())

	return c
}

func newGormMigrateSync() *cobra.Command {
	c := &cobra.Command{
		Use:   "sync",
		Short: "sync migrations code",
	}

	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run gorm cli")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		cfg := config.MustInit(env, false)

		a := app.New(cfg)
		m := migrations.New(a)
		m.Sync()
	}

	return c
}

func newGormMigrateCreate() *cobra.Command {
	c := &cobra.Command{
		Use:   "create",
		Short: "sync migrations code",
	}
	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run gorm cli")
	c.Flags().StringP(ArgVersion, ArgVersionShort, time.Now().Format("200601021504"), "set version to create migration")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		version, err := cmd.Flags().GetString(ArgVersion)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		cfg := config.MustInit(env, false)
		a := app.New(cfg)
		m := migrations.New(a)
		m.Create(version)
	}

	return c
}
