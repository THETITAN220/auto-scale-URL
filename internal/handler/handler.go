package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/THETITAN220/auto-scale-URL/internal/service"
)

type Handler struct {
	service *service.URLService
}

func NewHandler(s *service.URLService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing URL", http.StatusBadRequest)
		return
	}

	code := h.service.ShortenURL(url)
	shortURL := fmt.Sprintf("http://localhost:8080/r/%s", code)

	fmt.Fprintln(w, shortURL)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	code := strings.TrimPrefix(r.URL.Path, "/r/")
	original, ok := h.service.GetOriginalURL(code)

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, original, http.StatusFound)
}