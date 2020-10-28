package service

import (
	"golang-clean-architecture/entity"
	"golang-clean-architecture/model"
	"golang-clean-architecture/repository"
	"golang-clean-architecture/validation"
)

func NewProductService(productRepository *repository.ProductRepository) ProductService {
	return &productServiceImpl{
		ProductRepository: *productRepository,
	}
}

type productServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func (service *productServiceImpl) Create(request model.CreateProductRequest) (response model.CreateProductResponse) {
	validation.Validate(request)

	product := entity.Product{
		Id:       request.Id,
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	service.ProductRepository.Insert(product)

	response = model.CreateProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}
	return response
}

func (service *productServiceImpl) List() (responses []model.GetProductResponse) {
	products := service.ProductRepository.FindAll()
	for _, product := range products {
		responses = append(responses, model.GetProductResponse{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}
	return responses
}
