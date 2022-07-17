package services

import (
	"context"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sabilimaulana/sebelbucks-product-service/pkg/db"
	"github.com/sabilimaulana/sebelbucks-product-service/pkg/models"
	"github.com/sabilimaulana/sebelbucks-product-service/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *Server) ListProduct(ctx context.Context, _ *empty.Empty) (*pb.ListProductResponse, error) {
	productsDAO := []models.Product{}

	if result := s.H.DB.Find(&productsDAO); result.Error != nil {
		return &pb.ListProductResponse{
			Status:   http.StatusInternalServerError,
			Error:    result.Error.Error(),
			Products: []*pb.Product{},
		}, result.Error
	}

	products := []*pb.Product{}

	for _, product := range productsDAO {
		products = append(products, &pb.Product{
			Id:          int64(product.Id),
			Name:        product.Name,
			Description: product.Description,
			Stock:       product.Stock,
			Price:       product.Price,
			CreatedAt:   timestamppb.New(product.CreatedAt),
		})
	}

	return &pb.ListProductResponse{
		Status:   http.StatusOK,
		Products: products,
	}, nil
}
