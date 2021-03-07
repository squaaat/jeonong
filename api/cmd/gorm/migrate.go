package gorm

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/squaaat/nearsfeed/api/internal/config"
	_const "github.com/squaaat/nearsfeed/api/internal/const"
	"github.com/squaaat/nearsfeed/api/internal/container"
	"github.com/squaaat/nearsfeed/api/internal/db"
	categoryStore "github.com/squaaat/nearsfeed/api/internal/service/category/store"
	manufactureStore "github.com/squaaat/nearsfeed/api/internal/service/manufacture/store"
	"github.com/squaaat/nearsfeed/api/migrations"
)

func newGormMigrate() *cobra.Command {
	c := &cobra.Command{
		Use:   "migrate",
		Short: "it's about gorm migrate",
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

			catStore := categoryStore.New(cont)
			err = catStore.MustLoadDataAtLocal()
			if err != nil {
				return err
			}

			manStore := manufactureStore.New(cont)
			err = manStore.MustLoadDataAtLocal()
			if err != nil {
				return err
			}

			log.Info().Msg("migrating schema and loading data are done")
			return nil
		},
	}

	c.Flags().String(_const.ArgEnv, _const.ArgEnvDefault, "set environment to run gorm cli")

	return c
}

func newGormMigrateCreate() *cobra.Command {
	c := &cobra.Command{
		Use:   "create",
		Short: "sync migrations code",
		RunE: func(cmd *cobra.Command, _ []string) error {
			env, err := cmd.Flags().GetString(_const.ArgEnv)
			if err != nil {
				return err
			}

			version, err := cmd.Flags().GetString(_const.ArgVersion)
			if err != nil {
				return err
			}

			cfg, err := config.New(env)
			if err != nil {
				return err
			}
			cont, err := container.New(cfg)

			err = migrations.New(cont).Create(version)
			if err != nil {
				return err
			}

			log.Info().Msg("create migration sync template")
			return nil
		},
	}
	c.Flags().String(_const.ArgEnv, _const.ArgEnvDefault, "set environment to run gorm cli")
	c.Flags().String(_const.ArgVersion, time.Now().Format("200601021504"), "set version to create migration")

	return c
}
