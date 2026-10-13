package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/local/req/internal/profile"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show <name>",
	Short: "Show a profile's configuration",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := profile.Get(args[0])
		if err != nil {
			return err
		}
		out, _ := json.MarshalIndent(p, "", "  ")
		fmt.Println(string(out))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
