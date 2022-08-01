package service

import (
	"context"
	"position_service/config"
	pb "position_service/genproto/position_service"
	"position_service/pkg/logger"
	"position_service/storage"
)

type positionService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedPositionServiceServer
}

func NewPositionService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *positionService {
	return &positionService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *positionService) CreatePosition(ctx context.Context, req *pb.CreatePositionRequest) (*pb.PositionId, error) {
	id, err := s.strg.Position().Create(ctx, req)
	if err != nil {
		s.log.Error("CreateAttribute", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.PositionId{
		Id: id.Id,
	}, nil

}

func (s *positionService) GetAllPosition(ctx context.Context, req *pb.GetAllPositionRequest) (*pb.GetAllPositionResponse, error) {
	resp, err := s.strg.Position().GetAll(ctx, req)
	if err != nil {
		s.log.Error("GetAllPosition", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *positionService) UpdatePosition(ctx context.Context, req *pb.UpdatePositionRequest) (*pb.PositionId, error) {

	position_id, err := s.strg.Position().Update(ctx, req)
	if err != nil {
		s.log.Error("UpdatePosition", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return position_id, nil
}

func (s *positionService) DeletePosition(ctx context.Context, req *pb.PositionId) (*pb.PositionId, error) {
	position_id, err := s.strg.Position().Delete(ctx, req)
	if err != nil {
		s.log.Error("DeletePosition", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return position_id, nil
}

func (s *positionService) GetByIdPosition(ctx context.Context, req *pb.PositionId) (*pb.Position, error) {
	resp, err := s.strg.Position().GetById(ctx, req)
	if err != nil {
		s.log.Error("DeletePosition", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}
