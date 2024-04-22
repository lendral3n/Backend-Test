package service

import (
	"errors"
	"lendra/features/product"
	"lendra/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.ProductDataInterface)
	srv := New(repo)

	returnData := product.Product{
		ID:          1,
		NameProduct: "samsung",
		Price:       100000,
		UserID:      1,
	}

	t.Run("invalid name produk", func(t *testing.T) {
		caseData := returnData
		nameProduk := caseData.NameProduct
		caseData.NameProduct = ""
		err := srv.Create(1, caseData)

		assert.Error(t, err)
		assert.Equal(t, "nama produk tidak boleh kosong", err.Error())

		caseData.NameProduct = nameProduk
	})

	t.Run("invalid price produk", func(t *testing.T) {
		caseData := returnData
		priceProduk := caseData.Price
		caseData.Price = 0
		err := srv.Create(1, caseData)

		assert.Error(t, err)
		assert.Equal(t, "harga produk harus lebih besar dari 0", err.Error())
		caseData.Price = priceProduk
	})

	t.Run("error from repository", func(t *testing.T) {
		repo.On("Create", 1, returnData).Return(errors.New("database error")).Once()

		err := srv.Create(1, returnData)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseData := returnData
		repo.On("Create", 1, caseData).Return(nil).Once()

		err := srv.Create(1, caseData)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockRepo := new(mocks.ProductDataInterface)
	productService := New(mockRepo)

	t.Run("error from repository", func(t *testing.T) {
		mockRepo.On("Delete", 1, 1).Return(errors.New("database error")).Once()

		err := productService.Delete(1, 1)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Delete", 1, 1).Return(nil).Once()

		err := productService.Delete(1, 1)

		assert.NoError(t, err)
	})
}

func TestGetAllProduct(t *testing.T) {
	mockRepo := new(mocks.ProductDataInterface)
	productService := New(mockRepo)

	t.Run("error from repository", func(t *testing.T) {
		mockRepo.On("GetAllProduct").Return(nil, errors.New("database error")).Once()

		_, err := productService.GetAllProduct()

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetAllProduct").Return([]product.Product{}, nil).Once()

		_, err := productService.GetAllProduct()

		assert.NoError(t, err)
	})
}

func TestGetProductById(t *testing.T) {
	mockRepo := new(mocks.ProductDataInterface)
	productService := New(mockRepo)

	t.Run("error from repository", func(t *testing.T) {
		mockRepo.On("GetProductById", 1).Return(nil, errors.New("database error")).Once()

		_, err := productService.GetProductById(1)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetProductById", 1).Return(&product.Product{}, nil).Once()

		_, err := productService.GetProductById(1)

		assert.NoError(t, err)
	})
}

func TestUpdate(t *testing.T) {
	mockRepo := new(mocks.ProductDataInterface)
	productService := New(mockRepo)

	t.Run("error from repository", func(t *testing.T) {
		mockRepo.On("Update", 1, 1, mock.Anything).Return(errors.New("database error")).Once()

		err := productService.Update(1, 1, product.Product{})

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Update", 1, 1, mock.Anything).Return(nil).Once()

		err := productService.Update(1, 1, product.Product{})

		assert.NoError(t, err)
	})
}
