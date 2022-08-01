package service

import (
	"context"
	"position_service/config"
	pb "position_service/genproto/position_service"
	"position_service/pkg/logger"
	"position_service/storage"
)

type attributeService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedAttributeServiceServer
}

func NewAttributeService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *attributeService {
	return &attributeService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *attributeService) CreateAttribute(ctx context.Context, req *pb.CreateAttributeRequest) (*pb.Attribute, error) {
	id, err := s.strg.Attribute().CreateAttribute(ctx, req)
	if err != nil {
		s.log.Error("CreateAttribute", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.Attribute{
		Id:            id,
		Name:          req.Name,
		AttributeType: req.AttributeType,
	}, nil
}

func (s *attributeService) UpdateAttribute(ctx context.Context, req *pb.Attribute) (*pb.Attribute, error) {

	givenId, err := s.strg.Attribute().UpdateAttribute(ctx, req)
	if err != nil {
		s.log.Error("UpdateAttribute", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.Attribute{
		Id:            givenId,
		Name:          req.Name,
		AttributeType: req.AttributeType,
	}, nil

}

func (s *attributeService) GetAllAttribute(ctx context.Context, req *pb.GetAllAttributeRequest) (*pb.GetAllAttributeResponse, error) {
	resp, err := s.strg.Attribute().GetAllAttribute(ctx, req)
	if err != nil {
		s.log.Error("GetAllAttribute", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}
func (s *attributeService) GetByIdAttribute(ctx context.Context, req *pb.AttributeId) (*pb.Attribute, error) {
	resp, err := s.strg.Attribute().GetByIdAttribute(ctx, req)
	if err != nil {
		s.log.Error("GetByIdAttribute", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *attributeService) DeleteAttribute(ctx context.Context, req *pb.AttributeId) (*pb.AttributeId, error) {
	id, err := s.strg.Attribute().DeleteAttribute(ctx, req)
	if err != nil {
		s.log.Error("DeleteAttribute", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.AttributeId{
		Id: id,
	}, nil
}
