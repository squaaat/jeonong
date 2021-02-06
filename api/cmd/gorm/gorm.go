package gorm

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/squaaat/jeonong/api/internal/app"
	"github.com/squaaat/jeonong/api/internal/config"
	"github.com/squaaat/jeonong/api/internal/db"
	"github.com/squaaat/jeonong/api/migrations"
)

const (
	ArgEnv        = "environment"
	ArgEnvShort   = "e"
	ArgEnvDefault = "alpha"

	ArgVersion      = "version"
	ArgVersionShort = "v"
)

func New() *cobra.Command {
	c := &cobra.Command{
		Use:     "gorm",
		Short:   "jeonong-api cli gorm scripts",
		Aliases: []string{"g"},
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
	c.AddCommand(newGormClean())
	c.AddCommand(newGormCreate())
	c.AddCommand(newGormReCreate())
	c.AddCommand(newGormMigrate())

	return c
}

func newGormReCreate() *cobra.Command {
	c := &cobra.Command{
		Use:   "re-create",
		Short: "create schema for develop",
	}
	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run gorm cli")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		cfg := config.MustInit(env, false)

		a := app.New(cfg)
		a.ServiceDB.Clean()

		err = db.Initialize(cfg.ServiceDB)

		m := migrations.New(a)
		m.Sync()

		if err != nil {
			log.Fatal().Err(err).Send()
		}
	}

	return c
}

func newGormCreate() *cobra.Command {
	c := &cobra.Command{
		Use:   "create",
		Short: "create schema for develop",
	}
	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run gorm cli")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		cfg := config.MustInit(env, false)

		err = db.Initialize(cfg.ServiceDB)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		a := app.New(cfg)
		m := migrations.New(a)
		m.Sync()

		if err != nil {
			log.Fatal().Err(err).Send()
		}
	}

	return c
}

func newGormClean() *cobra.Command {
	c := &cobra.Command{
		Use:   "clean",
		Short: "remove schema(db)",
	}
	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run gorm cli")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		cfg := config.MustInit(env, false)

		a := app.New(cfg)
		a.ServiceDB.Clean()
	}

	return c
}
