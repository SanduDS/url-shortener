package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SanduDS/url-shortener/internal/model"
	"github.com/SanduDS/url-shortener/internal/service"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(s *service.URLService) *URLHandler {
	return &URLHandler{
		service: s,
	}
}

func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req model.URLRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortCode := h.service.ShortenURL(req.URL)
	baseURL := "http://localhost:8080/"
	resp := model.URLResponse{
		ShortURL: baseURL + shortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
		return
	}

}
