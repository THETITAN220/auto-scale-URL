package service

import (
	"github.com/THETITAN220/auto-scale-URL/internal/storage"
	"github.com/THETITAN220/auto-scale-URL/pkg/utils"
)

type URLService struct {
	store storage.Store
}

func NewURLService(s storage.Store) *URLService {
	return &URLService{store: s}
}

func (s *URLService) ShortenURL(url string) string {
	code := utils.GenerateCode()
	s.store.Save(code, url)
	return code
}

func (s *URLService) GetOriginalURL(code string) (string, bool) {
	return s.store.Get(code)
}