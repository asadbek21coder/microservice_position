package service

import (
	"context"
	"position_service/config"

	pb "position_service/genproto/position_service"
	"position_service/pkg/logger"
	"position_service/storage"
)

type professionService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedProfessionServiceServer
}

func NewProfessionService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *professionService {
	return &professionService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *professionService) Create(ctx context.Context, req *pb.CreateProfessionRequest) (*pb.Profession, error) {
	id, err := s.strg.Profession().Create(ctx, req)
	if err != nil {
		s.log.Error("CreateProfession", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.Profession{
		Id:   id,
		Name: req.Name,
	}, nil
}

func (s *professionService) GetAll(ctx context.Context, req *pb.GetAllProfessionRequest) (*pb.GetAllProfessionResponse, error) {
	resp, err := s.strg.Profession().GetAll(ctx, req)
	if err != nil {
		s.log.Error("GetAllProfession", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *professionService) Update(ctx context.Context, req *pb.Profession) (*pb.Profession, error) {

	givenId, err := s.strg.Profession().Update(ctx, req)
	if err != nil {
		s.log.Error("UpdateProfession", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.Profession{
		Id:   givenId,
		Name: req.Name,
	}, nil
}

func (s *professionService) Delete(ctx context.Context, req *pb.Id) (*pb.Id, error) {
	id, err := s.strg.Profession().Delete(ctx, req)
	if err != nil {
		s.log.Error("DeleteProfession", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.Id{
		Id: id,
	}, nil
}

func (s *professionService) GetById(ctx context.Context, req *pb.Id) (*pb.Profession, error) {
	resp, err := s.strg.Profession().GetById(ctx, req)
	if err != nil {
		s.log.Error("DeleteProfession", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}
