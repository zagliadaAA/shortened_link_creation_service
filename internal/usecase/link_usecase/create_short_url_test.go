package link_usecase

import (
	"errors"
	"testing"

	"shortened_link_creation_service/internal/domain"
	"shortened_link_creation_service/internal/usecase/link_usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateShortURLUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("error test")

	//зависимости, которые нужны для теста
	type fields struct {
		linkRepo *mocks.Repo
	}

	//данные для теста
	type args struct {
		URL string
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
				URL: "https://google.com",
			},
			want: &domain.Link{
				URL: "https://google.com",
			},
			before: func(f fields, args args) {
				// mack.Anything из-за случайной генерации shortURL
				f.linkRepo.EXPECT().Create(mock.Anything).Return(&domain.Link{URL: args.URL}, nil)
			},
		},
		{
			name: "error on create link",
			args: args{
				URL: "https://google.com",
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				f.linkRepo.EXPECT().Create(mock.Anything).Return(nil, errTest)
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

			//выполняем
			result, err := uc.CreateShortURL(tt.args.URL)

			//проверяем результат
			if tt.wantErr != nil {
				a.ErrorIs(err, tt.wantErr)
				return
			}

			//проверяем только url, shortURL создается случайно
			a.NoError(err)
			if tt.want != nil {
				a.Equal(tt.want.URL, result.URL)
			}
		})
	}

}
