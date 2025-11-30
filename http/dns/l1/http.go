package l1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIPAddress(domain string) (string, error) {
	url := fmt.Sprintf("https://1.1.1.1/dns-query?name=%s&type=A", domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/dns-json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var response DNSResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}

	// Find A record
	for _, answer := range response.Answer {
		if answer.Type == 1 { // Type 1 = A record
			return answer.Data, nil // Return the IP address
		}
	}

	return "", fmt.Errorf("no A record found for %s", domain)
}
	