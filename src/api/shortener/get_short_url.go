package shortener

import (
	"context"

	desc "github.com/kaspeya/go-url-shortener/pkg/shortener"
)

func (i *Implementation) GetShortUrl(ctx context.Context, req *desc.GetShortUrlRequest) (*desc.GetShortUrlResponse, error) {
	shortUrl, err := i.shortenerService.GetShortUrl(ctx, req.GetOriginalUrl())
	if err != nil {
		return nil, err
	}

	return &desc.GetShortUrlResponse{
		ShortUrl: shortUrl,
	}, nil
}
