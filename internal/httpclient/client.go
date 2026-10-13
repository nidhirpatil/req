package httpclient

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/local/req/internal/profile"
)

type Response struct {
	StatusCode int
	Status     string
	Headers    http.Header
	Body       string
	Duration   time.Duration
}

func Execute(p profile.Profile) (*Response, error) {
	method := p.Method
	if method == "" {
		method = "GET"
	}

	var bodyReader io.Reader
	if p.Body != "" {
		bodyReader = strings.NewReader(p.Body)
	}

	req, err := http.NewRequest(method, p.URL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	for k, v := range p.Headers {
		req.Header.Set(k, v)
	}

	switch p.Auth.Type {
	case "bearer":
		req.Header.Set("Authorization", "Bearer "+p.Auth.Token)
	case "basic":
		creds := base64.StdEncoding.EncodeToString([]byte(p.Auth.User + ":" + p.Auth.Pass))
		req.Header.Set("Authorization", "Basic "+creds)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Headers:    resp.Header,
		Body:       string(body),
		Duration:   elapsed,
	}, nil
}
