package categories

type usecase struct {
	repo Repository
}

type Usecase interface {
	getCategories() ([]Category, error)
	createCategory(Category) (Category, error)
	updateCategory(string, Category) (Category, error)
}

func NewUsecase(repo Repository) Usecase {
	return usecase{repo}
}

func (usecase usecase) getCategories() ([]Category, error) {
	return usecase.repo.getCategories()
}

func (usecase usecase) createCategory(category Category) (Category, error) {
	return usecase.repo.createCategory(category)
}

func (usecase usecase) updateCategory(serial string, category Category) (Category, error) {
	return usecase.repo.updateCategory(serial, category)
}
