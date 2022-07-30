package services

import (
	"context"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sabilimaulana/sebelbucks-product-service/pkg/models"
	"github.com/sabilimaulana/sebelbucks-product-service/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *Server) ListVariant(ctx context.Context, _ *empty.Empty) (*pb.ListVariantResponse, error) {
	variantsDAO := []models.Variant{}

	if result := s.H.DB.Find(&variantsDAO); result.Error != nil {
		return &pb.ListVariantResponse{
			Status:   http.StatusInternalServerError,
			Error:    result.Error.Error(),
			Variants: []*pb.Variant{},
		}, result.Error
	}

	variants := []*pb.Variant{}

	// Formatter
	// TODO: refractor this to utils
	for _, variant := range variantsDAO {
		variants = append(variants, &pb.Variant{
			Id:          int64(variant.Id),
			Name:        variant.Name,
			Description: variant.Description,
			Stock:       variant.Stock,
			Price:       variant.Price,
			CreatedAt:   timestamppb.New(variant.CreatedAt),
			UpdatedAt:   timestamppb.New(variant.UpdatedAt),
		})
	}

	return &pb.ListVariantResponse{
		Status:   http.StatusOK,
		Variants: variants,
	}, nil
}
