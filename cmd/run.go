package cmd

import (
	"github.com/local/req/internal/display"
	"github.com/local/req/internal/httpclient"
	"github.com/local/req/internal/profile"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <name>",
	Short: "Execute a saved request profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := profile.Get(args[0])
		if err != nil {
			return err
		}
		resp, err := httpclient.Execute(p)
		if err != nil {
			display.PrintError(err)
			return nil
		}
		display.PrintResponse(resp)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
