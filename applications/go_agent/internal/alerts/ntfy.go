package alerts

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Send(topic, title, message string) error {
	if topic == "" {
		return fmt.Errorf("NTFY_TOPIC is not configured")
	}

	url := "https://ntfy.sh/" + topic

	req, err := http.NewRequest(
		http.MethodPost,
		url,
		strings.NewReader(message),
	)
	if err != nil {
		return fmt.Errorf("create ntfy request: %w", err)
	}

	req.Header.Set("Title", title)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("send ntfy notification: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		body, _ := io.ReadAll(response.Body)

		retun fmt.Errorf(
			"ntfy returned %s: %s",
			response.Status,
			string(body),
		)
	}

	return nil
}
