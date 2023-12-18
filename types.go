package spn2

// CaptureResponse represents the response from a capture request.
type CaptureResponse struct {
	URL   string `json:"url"`
	JobID string `json:"job_id"`
}

// StatusResponse represents the response from a status request.
type StatusResponse struct {
	Counters struct {
		Embeds   int `json:"embeds"`
		Outlinks int `json:"outlinks"`
	} `json:"counters"`
	DurationSec  float64  `json:"duration_sec"`
	FirstArchive bool     `json:"first_archive"`
	HTTPStatus   int      `json:"http_status"`
	JobID        string   `json:"job_id"`
	OriginalURL  string   `json:"original_url"`
	Outlinks     []string `json:"outlinks"`
	Resources    []string `json:"resources"`
	Status       string   `json:"status"`
	Timestamp    string   `json:"timestamp"`
}

// SystemStatusResponse represents the system status response.
type SystemStatusResponse struct {
	Status string `json:"status"`
}

// UserStatusResponse represents the user status response.
type UserStatusResponse struct {
	Available  int `json:"available"`
	Processing int `json:"processing"`
}
