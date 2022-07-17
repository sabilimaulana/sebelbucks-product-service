package services

import (
	"context"
	"net/http"

	"github.com/sabilimaulana/sebelbucks-product-service/pkg/db"
	"github.com/sabilimaulana/sebelbucks-product-service/pkg/models"
	"github.com/sabilimaulana/sebelbucks-product-service/pkg/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedProductServiceServer
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product

	product.Name = req.Name
	product.Stock = req.Stock
	product.Price = req.Price
	product.Description = req.Description

	if result := s.H.DB.Create(&product); result.Error != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateProductResponse{
		Status:    http.StatusCreated,
		ProductId: int64(product.Id),
	}, nil
}
