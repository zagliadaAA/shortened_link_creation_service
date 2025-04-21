package link_usecase

type UseCase struct {
	repo repo
}

type repo interface {
	//функции у 2 репозиториев должны быть одни и те же
}

func NewUseCase(repo repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
