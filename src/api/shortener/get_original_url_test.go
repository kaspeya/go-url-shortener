package shortener

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	desc "github.com/kaspeya/go-url-shortener/pkg/shortener"
	shortenerMocks "github.com/kaspeya/go-url-shortener/src/repository/shortener/mocks"
	"github.com/kaspeya/go-url-shortener/src/service/shortener"
	"github.com/stretchr/testify/require"
)

func TestGetOriginalUrl(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		shortUrl    = "https://shorturl.com/hjhjgjhlfgl"
		originalUrl = gofakeit.URL()
		repoErrText = gofakeit.Phrase()

		req = &desc.GetOriginalUrlRequest{
			ShortUrl: shortUrl,
		}

		validRes = &desc.GetOriginalUrlResponse{
			OriginalUrl: originalUrl,
		}

		repoErr = errors.New(repoErrText)
	)

	shortenerMock := shortenerMocks.NewMockRepository(mockCtrl)
	gomock.InOrder(
		shortenerMock.EXPECT().GetOriginalUrl(ctx, gomock.Any()).Return(originalUrl, nil),
		shortenerMock.EXPECT().GetOriginalUrl(ctx, gomock.Any()).Return("", repoErr),
	)

	api := newMockImplementation(Implementation{
		shortenerService: shortener.NewServiceMock(shortenerMock),
	})

	t.Run("success case", func(t *testing.T) {
		res, err := api.GetOriginalUrl(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("repo err", func(t *testing.T) {
		_, err := api.GetOriginalUrl(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
