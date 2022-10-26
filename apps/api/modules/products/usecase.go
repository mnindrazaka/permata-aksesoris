package products

type usecase struct {
	repo Repository
}

type Usecase interface {
	getProducts() ([]Product, error)
}

func NewUsecase(repo Repository) Usecase {
	return usecase{repo}
}

func (usecase usecase) getProducts() ([]Product, error) {
	return usecase.repo.getProducts()
}
