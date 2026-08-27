package alerts

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	ntfyBaseURL    = "https://ntfy.sh/"
	requestTimeout = 10 * time.Second
	maxErrorBody   = 4096
)

var client = &http.Client{
	Timeout: requestTimeout,
}

func Send(topic, title, message string) error {
	topic = strings.TrimSpace(topic)
	title = strings.TrimSpace(title)

	if topic == "" {
		return fmt.Errorf("NTFY_TOPIC is not configured")
	}

	if message == "" {
		return fmt.Errorf("notification message is empty")
	}

	topicURL := ntfyBaseURL + url.PathEscape(topic)

	req, err := http.NewRequest(
		http.MethodPost,
		topicURL,
		strings.NewReader(message),
	)
	if err != nil {
		return fmt.Errorf("create ntfy request: %w", err)
	}

	if title != "" {
		req.Header.Set("Title", title)
	}

	req.Header.Set("Content-Type", "text/plain; charset=utf-8")

	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("send ntfy notification: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode < http.StatusOK ||
		response.StatusCode >= http.StatusMultipleChoices {

		body, readErr := io.ReadAll(
			io.LimitReader(response.Body, maxErrorBody),
		)
		if readErr != nil {
			return fmt.Errorf(
				"ntfy returned %s and response body could not be read: %w",
				response.Status,
				readErr,
			)
		}

		responseMessage := strings.TrimSpace(string(body))

		if responseMessage == "" {
			return fmt.Errorf("ntfy returned %s", response.Status)
		}

		return fmt.Errorf(
			"ntfy returned %s: %s",
			response.Status,
			responseMessage,
		)
	}

	return nil
}
