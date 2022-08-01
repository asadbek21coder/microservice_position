package service

import (
	"context"
	"position_service/config"
	pb "position_service/genproto/company_service"
	"position_service/pkg/logger"
	"position_service/storage"
)

type companyService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedCompanyServiceServer
}

func NewCompanyeService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *companyService {
	return &companyService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *companyService) Create(ctx context.Context, req *pb.CreateCompanyRequest) (*pb.Company, error) {
	id, err := s.strg.Company().Create(ctx, req)
	if err != nil {
		s.log.Error("CreateAttribute", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.Company{
		Id:   id,
		Name: req.Name,
	}, nil
}

func (s *companyService) GetAll(ctx context.Context, req *pb.GetAllCompanyRequest) (*pb.GetAllCompanyResponse, error) {
	resp, err := s.strg.Company().GetAll(ctx, req)
	if err != nil {
		s.log.Error("GetAllProfession", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *companyService) Update(ctx context.Context, req *pb.Company) (*pb.Company, error) {

	givenId, err := s.strg.Company().Update(ctx, req)
	if err != nil {
		s.log.Error("UpdateProfession", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.Company{
		Id:   givenId,
		Name: req.Name,
	}, nil
}

func (s *companyService) Delete(ctx context.Context, req *pb.Id) (*pb.Id, error) {
	id, err := s.strg.Company().Delete(ctx, req)
	if err != nil {
		s.log.Error("DeleteProfession", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.Id{
		Id: id,
	}, nil
}

func (s *companyService) GetById(ctx context.Context, req *pb.Id) (*pb.Company, error) {
	resp, err := s.strg.Company().GetById(ctx, req)
	if err != nil {
		s.log.Error("DeleteProfession", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}
