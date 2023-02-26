package command

import (
	"errors"
	"github.com/gitscan/config"
	"github.com/gitscan/internal/database"
	"github.com/gitscan/internal/migrations"
	"github.com/golang-migrate/migrate/v4"

	"github.com/spf13/cobra"
)

func configureDbCommand(command *cobra.Command) {
	rootCommand := &cobra.Command{
		Use:   "db",
		Short: "manipulate database",
	}
	migrateCommand := &cobra.Command{
		Use:   "migrate",
		Short: "apply sql",
		RunE:  dbMigrate,
	}
	migrateCommand.PersistentFlags().Bool("no-lock", false, "use --no-lock to disable locking")
	dropCommand := &cobra.Command{
		Use:   "drop",
		Short: "drop everything in db",
		RunE:  dbDrop,
	}
	dropCommand.PersistentFlags().Bool("no-lock", false, "use --no-lock to disable locking")
	command.AddCommand(rootCommand)
	rootCommand.AddCommand(migrateCommand)
	rootCommand.AddCommand(dropCommand)
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	noLock, err := cmd.Flags().GetBool("no-lock")
	if err != nil {
		noLock = false
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := database.New(cfg.DB)
	if err != nil {
		panic(err)
	}
	migration, err := migrations.New(db.RawDB(), noLock)

	if err != nil {
		return err
	}
	cmd.Println("Migrating...")
	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func dbDrop(cmd *cobra.Command, args []string) error {
	noLock, err := cmd.Flags().GetBool("no-lock")
	if err != nil {
		noLock = false
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := database.New(cfg.DB)
	if err != nil {
		panic(err)
	}

	migration, err := migrations.New(db.RawDB(), noLock)
	if err != nil {
		return err
	}

	cmd.Println("Deleting...")

	return migration.Drop()
}
