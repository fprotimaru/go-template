package cmd

import (
	"github.com/spf13/cobra"
)

type commands struct {
	addr, cfg string
}

var cmd commands

func init() {
	rootCmd.PersistentFlags().StringVarP(&cmd.cfg, "conf", "c", "config.yaml", "config file")
	rootCmd.PersistentFlags().StringVarP(&cmd.addr, "addr", "a", ":8000", "address and port")
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	rootCmd.CompletionOptions = cobra.CompletionOptions{DisableDefaultCmd: true}
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(createUserCmd)
}

var rootCmd = &cobra.Command{
	Use: "app-name",
}

func Execute() error {
	return rootCmd.Execute()
}
