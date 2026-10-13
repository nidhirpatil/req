package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/local/req/internal/profile"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		profiles, err := profile.LoadAll()
		if err != nil {
			return err
		}
		if len(profiles) == 0 {
			fmt.Println("No profiles saved yet. Use `req save` to create one.")
			return nil
		}
		bold := color.New(color.Bold)
		dim := color.New(color.FgHiBlack)
		for _, p := range profiles {
			bold.Printf("  %-20s", p.Name)
			dim.Printf("  %s %s\n", p.Method, p.URL)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
