package server

import (
	"github.com/spf13/cobra"
)


func init() {
}

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Start an gps api http server used by the mobile clients",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

