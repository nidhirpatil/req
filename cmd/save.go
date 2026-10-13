package cmd

import (
	"fmt"
	"strings"

	"github.com/local/req/internal/profile"
	"github.com/spf13/cobra"
)

var (
	saveMethod  string
	saveURL     string
	saveHeaders []string
	saveBody    string
	saveBearer  string
	saveUser    string
	savePass    string
)

var saveCmd = &cobra.Command{
	Use:   "save <name>",
	Short: "Save a request profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		headers := make(map[string]string)
		for _, h := range saveHeaders {
			parts := strings.SplitN(h, ":", 2)
			if len(parts) != 2 {
				return fmt.Errorf("invalid header %q — use Key:Value format", h)
			}
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}

		p := profile.Profile{
			Name:    name,
			Method:  strings.ToUpper(saveMethod),
			URL:     saveURL,
			Headers: headers,
			Body:    saveBody,
		}

		if saveBearer != "" {
			p.Auth = profile.Auth{Type: "bearer", Token: saveBearer}
		} else if saveUser != "" {
			p.Auth = profile.Auth{Type: "basic", User: saveUser, Pass: savePass}
		}

		if err := profile.Save(p); err != nil {
			return err
		}
		fmt.Printf("Saved profile %q\n", name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
	saveCmd.Flags().StringVarP(&saveMethod, "method", "X", "GET", "HTTP method")
	saveCmd.Flags().StringVarP(&saveURL, "url", "u", "", "Request URL")
	saveCmd.Flags().StringArrayVarP(&saveHeaders, "header", "H", nil, "Header in Key:Value format (repeatable)")
	saveCmd.Flags().StringVarP(&saveBody, "body", "d", "", "Request body")
	saveCmd.Flags().StringVar(&saveBearer, "bearer", "", "Bearer token for Authorization header")
	saveCmd.Flags().StringVar(&saveUser, "user", "", "Username for basic auth")
	saveCmd.Flags().StringVar(&savePass, "pass", "", "Password for basic auth")
	saveCmd.MarkFlagRequired("url")
}
