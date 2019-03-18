package cmd

import (
	"github.com/devchallenge/spy-api/cmd/server"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "spy-api",
	Short: "Spy API",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
	}
}

func init() {
	RootCmd.AddCommand(server.Cmd)
}
