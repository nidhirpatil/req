package display

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/local/req/internal/httpclient"
)

var (
	green  = color.New(color.FgGreen, color.Bold)
	red    = color.New(color.FgRed, color.Bold)
	yellow = color.New(color.FgYellow)
	dim    = color.New(color.FgHiBlack)
	bold   = color.New(color.Bold)
)

func PrintResponse(resp *httpclient.Response) {
	// Status line
	statusColor := green
	if resp.StatusCode >= 400 {
		statusColor = red
	} else if resp.StatusCode >= 300 {
		statusColor = yellow
	}
	statusColor.Printf("HTTP %s", resp.Status)
	dim.Printf("  (%s)\n\n", resp.Duration.Round(1))

	// Headers
	bold.Println("Headers:")
	for k, vals := range resp.Headers {
		for _, v := range vals {
			dim.Printf("  %s: ", k)
			fmt.Println(v)
		}
	}
	fmt.Println()

	// Body
	bold.Println("Body:")
	if isJSON(resp.Headers) {
		prettyJSON(resp.Body)
	} else {
		fmt.Println(resp.Body)
	}
}

func isJSON(h http.Header) bool {
	ct := h.Get("Content-Type")
	return len(ct) > 0 && (ct[:16] == "application/json" || containsJSON(ct))
}

func containsJSON(s string) bool {
	for i := 0; i+4 < len(s); i++ {
		if s[i:i+4] == "json" {
			return true
		}
	}
	return false
}

func prettyJSON(body string) {
	var v interface{}
	if err := json.Unmarshal([]byte(body), &v); err != nil {
		fmt.Println(body)
		return
	}
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(body)
		return
	}
	fmt.Println(string(out))
}

func PrintError(err error) {
	red.Fprintf(color.Output, "Error: %v\n", err)
}
