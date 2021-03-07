package gorm

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/squaaat/nearsfeed/api/internal/config"
	_const "github.com/squaaat/nearsfeed/api/internal/const"
	"github.com/squaaat/nearsfeed/api/internal/container"
	"github.com/squaaat/nearsfeed/api/internal/db"
	"github.com/squaaat/nearsfeed/api/migrations"
)

func New() *cobra.Command {
	c := &cobra.Command{
		Use:   "gorm",
		Short: "api cli gorm scripts",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
	c.AddCommand(newGormDropDB())
	c.AddCommand(newGormCreate())
	c.AddCommand(newGormReCreate())
	c.AddCommand(newGormMigrate())

	return c
}

func newGormReCreate() *cobra.Command {
	c := &cobra.Command{
		Use:   "re-create",
		Short: "create schema for develop",
		RunE: func(cmd *cobra.Command, _ []string) error {
			env, err := cmd.Flags().GetString(_const.ArgEnv)
			if err != nil {
				return err
			}

			cfg, err := config.New(env)
			if err != nil {
				return err
			}

			err = db.DropDB(env, cfg.ServiceDB)
			if err != nil {
				return err
			}

			err = db.CreateDB(env, cfg.ServiceDB)
			if err != nil {
				return err
			}

			cont, err := container.New(cfg)
			if err != nil {
				return err
			}

			err = migrations.New(cont).Sync()
			if err != nil {
				return err
			}
			log.Info().Msg("gorm dropDB, createDB and migration success")
			return nil
		},
	}
	c.Flags().String(_const.ArgEnv, _const.ArgEnvDefault, "set environment to run gorm cli")

	return c
}

func newGormCreate() *cobra.Command {
	c := &cobra.Command{
		Use:   "create",
		Short: "create schema for develop",
		RunE: func(cmd *cobra.Command, _ []string) error {
			env, err := cmd.Flags().GetString(_const.ArgEnv)
			if err != nil {
				return err
			}

			cfg, err := config.New(env)
			if err != nil {
				return err
			}

			err = db.CreateDB(env, cfg.ServiceDB)
			if err != nil {
				return err
			}

			cont, err := container.New(cfg)
			if err != nil {
				return err
			}

			err = migrations.New(cont).Sync()
			if err != nil {
				return err
			}
			log.Info().Msg("gorm createDB and migration success")
			return nil
		},
	}
	c.Flags().String(_const.ArgEnv, _const.ArgEnvDefault, "set environment to run gorm cli")

	return c
}

func newGormDropDB() *cobra.Command {
	c := &cobra.Command{
		Use:   "drop",
		Short: "remove schema(db)",
		RunE: func(cmd *cobra.Command, _ []string) error {
			env, err := cmd.Flags().GetString(_const.ArgEnv)
			if err != nil {
				return err
			}

			cfg, err := config.New(env)
			if err != nil {
				return err
			}

			err = db.DropDB(env, cfg.ServiceDB)
			if err != nil {
				return err
			}
			log.Info().Msg("gorm dropDB")
			return nil
		},
	}
	c.Flags().String(_const.ArgEnv, _const.ArgEnvDefault, "set environment to run gorm cli")

	return c
}
