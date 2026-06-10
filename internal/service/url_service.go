package service

import (
	"sync"

	"github.com/SanduDS/url-shortener/pkg/base62"
)

type URLService struct {
	mu      sync.Mutex
	store   map[string]string // shortCode -> originalURL
	counter int
}

func NewURLService() *URLService {
	return &URLService{
		store: make(map[string]string),
	}
}

func (s *URLService) ShortenURL(url string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.counter = s.counter + 1
	shortCode := base62.Encode(s.counter)
	s.store[shortCode] = url
	return shortCode
}

func (s *URLService) GetOriginalURL(shortCode string) (string, bool) {
	url, ok := s.store[shortCode]
	return url, ok
}
