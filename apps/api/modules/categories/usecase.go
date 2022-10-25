package categories

type usecase struct {
	repo Repository
}

type Usecase interface {
	getCategories() ([]Category, error)
}

func NewUsecase(repo Repository) Usecase {
	return usecase{repo}
}

func (usecase usecase) getCategories() ([]Category, error) {
	return usecase.repo.getCategories()
}
