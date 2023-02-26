package command

import (
	"github.com/gitscan/http"
	"github.com/spf13/cobra"
)

func configureServerCommand(command *cobra.Command) {
	rootCommand := &cobra.Command{
		Use:   "server",
		Short: "manipulate server",
	}
	serverStartCommand := &cobra.Command{
		Use:   "start",
		Short: "start server",
		RunE:  startServer,
	}
	command.AddCommand(rootCommand)
	rootCommand.AddCommand(serverStartCommand)
}

func startServer(cmd *cobra.Command, args []string) error {
	http.StartServer()

	return nil
}
