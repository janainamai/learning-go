package mocks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository é um mock da interface Repository
type MockRepository struct {
	mock.Mock
}

// Save é a implementação mock do método Save
func (m *MockRepository) Save(product Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func TestSaveProduct(t *testing.T) {
	// Criando uma instância do mock
	mockRepo := new(MockRepository)

	// Configurando o comportamento esperado
	validProduct := Product{ID: 1, Name: "Valid Product", Price: 10.0}
	invalidProduct := Product{ID: 0, Name: "", Price: 0}

	mockRepo.On("Save", validProduct).Return(nil)
	mockRepo.On("Save", invalidProduct).Return(fmt.Errorf("invalid product ID"))

	// Teste para produto com ID válido
	err := SaveProduct(mockRepo, validProduct)
	assert.Nil(t, err)

	// Teste para produto com ID inválido
	err = SaveProduct(mockRepo, invalidProduct)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid product ID", err.Error())

	// Verifica se os métodos esperados foram chamados
	mockRepo.AssertExpectations(t)
}
