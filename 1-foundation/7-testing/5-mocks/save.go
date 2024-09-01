package mocks

type Product struct {
	ID    int
	Name  string
	Price float64
}

// Repository define a interface para salvar produtos
type Repository interface {
	Save(product Product) error
}

// RealRepository é uma implementação real de Repository
type RealRepository struct{}

// Save salva o produto no banco de dados (simulado)
func (r *RealRepository) Save(product Product) error {
	// Lógica para salvar no banco de dados real
	return nil
}

// SaveProduct salva um produto usando um Repository
func SaveProduct(repo Repository, product Product) error {
	return repo.Save(product)
}
