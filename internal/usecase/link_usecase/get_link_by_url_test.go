package link_usecase

import (
	"errors"
	"testing"

	"shortened_link_creation_service/internal/domain"
	"shortened_link_creation_service/internal/usecase/link_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetLinkByURLUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("error test")

	//зависимости, которые нужны для теста
	type fields struct {
		linkRepo *mocks.Repo
	}

	//данные для теста
	type args struct {
		URL      string
		shortURL string
	}

	//тесты
	tests := []struct {
		name    string
		args    args
		want    *domain.Link
		wantErr error
		before  func(f fields, args args)
	}{
		{
			name: "success",
			args: args{
				URL:      "https://google.com",
				shortURL: "https://myDomain.ru/Ur3j9E5tAd",
			},
			want: &domain.Link{
				URL:      "https://google.com",
				ShortURL: "https://myDomain.ru/Ur3j9E5tAd",
			},
			before: func(f fields, args args) {
				link := domain.NewLink(args.URL, args.shortURL)
				f.linkRepo.EXPECT().GetLinkByURL(args.URL).Return(link, nil)
			},
		},
		{
			name: "error get link by url",
			args: args{
				URL: "https://myDomain.ru/Ur3j9E5tAd",
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				f.linkRepo.EXPECT().GetLinkByURL(args.URL).Return(nil, errTest)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			//создали зависимости
			f := fields{
				linkRepo: mocks.NewRepo(t),
			}
			tt.before(f, tt.args)

			uc := NewUseCase(f.linkRepo)

			//выполнили
			e, err := uc.GetLinkByURL(tt.args.URL)

			//проверяем результат
			if tt.wantErr != nil {
				a.ErrorIs(err, tt.wantErr)
				return
			}

			a.NoError(err)
			a.Equal(tt.want, e)
		})
	}
}
