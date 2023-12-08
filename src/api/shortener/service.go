package shortener

import (
	desc "github.com/kaspeya/go-url-shortener/pkg/shortener"
	shortenerService "github.com/kaspeya/go-url-shortener/src/service/shortener"
)

type Implementation struct {
	desc.UnimplementedShortenerServer

	shortenerService shortenerService.Service
}

func NewImplementation(shortenerService shortenerService.Service) *Implementation {
	return &Implementation{
		shortenerService: shortenerService,
	}
}

func newMockImplementation(i Implementation) *Implementation {
	return &Implementation{
		desc.UnimplementedShortenerServer{},
		i.shortenerService,
	}
}
