package command

import (
	"github.com/spf13/cobra"
	"guardrail/gitscan/http"
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
	cmd.Println(http.StartServer())

	return nil
}
