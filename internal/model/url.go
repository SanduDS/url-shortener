package model

import "time"

type URLRequest struct {
	URL string `json:"url"`
}

type URLResponse struct {
	ShortURL string `json:"short_url"`
}

type URL struct {
	ID          string
	OriginalURL string
	ShortCode   string
	CreatedAt   time.Time
}
