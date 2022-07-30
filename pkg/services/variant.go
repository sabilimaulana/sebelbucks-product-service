package services

import (
	"context"
	"net/http"

	"github.com/sabilimaulana/sebelbucks-product-service/pkg/models"
	"github.com/sabilimaulana/sebelbucks-product-service/pkg/pb"
)

// Variant
func (s *Server) CreateVariant(ctx context.Context, req *pb.CreateVariantRequest) (*pb.CreateVariantResponse, error) {
	var variant models.Variant

	variant.Name = req.Name
	variant.Stock = req.Stock
	variant.Price = req.Price
	variant.Description = req.Description

	if result := s.H.DB.Create(&variant); result.Error != nil {
		return &pb.CreateVariantResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateVariantResponse{
		Status:    http.StatusCreated,
		VariantId: int64(variant.Id),
	}, nil
}
