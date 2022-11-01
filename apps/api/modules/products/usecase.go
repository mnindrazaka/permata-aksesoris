package products

type usecase struct {
	repo Repository
}

type Usecase interface {
	getProducts() ([]Product, error)
	getProductDetail(string) (Product, error)
	createProduct(Product) (Product, error)
	updateProduct(string, Product) (Product, error)
	deleteProduct(string) error

	createProductImage(ProductImage) (ProductImage, error)
	deleteProductImage(string) error
}

func NewUsecase(repo Repository) Usecase {
	return usecase{repo}
}

func (usecase usecase) getProducts() ([]Product, error) {
	return usecase.repo.getProducts()
}

func (usecase usecase) getProductDetail(serial string) (Product, error) {
	return usecase.repo.getProductDetail(serial)
}

func (usecase usecase) createProduct(product Product) (Product, error) {
	return usecase.repo.createProduct(product)
}

func (usecase usecase) updateProduct(serial string, product Product) (Product, error) {
	return usecase.repo.updateProduct(serial, product)
}

func (usecase usecase) deleteProduct(serial string) error {
	return usecase.repo.deleteProduct(serial)
}

func (usecase usecase) createProductImage(productImage ProductImage) (ProductImage, error) {
	return usecase.repo.createProductImage(productImage)
}

func (usecase usecase) deleteProductImage(serial string) error {
	return usecase.repo.deleteProductImage(serial)
}
