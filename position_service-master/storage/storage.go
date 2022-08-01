package storage

import (
	"context"
	"errors"
	pc "position_service/genproto/company_service"
	pb "position_service/genproto/position_service"
)

var ErrorTheSameId = errors.New("cannot use the same uuid for 'id' and 'parent_id' fields")
var ErrorProjectId = errors.New("not valid 'project_id'")

type StorageI interface {
	Profession() ProfessionI
	Attribute() AttributeI
	Company() CompanyI
	Position() PositionI
}

type ProfessionI interface {
	Create(ctx context.Context, entity *pb.CreateProfessionRequest) (id string, err error)
	GetAll(ctx context.Context, req *pb.GetAllProfessionRequest) (*pb.GetAllProfessionResponse, error)
	GetById(ctx context.Context, req *pb.Id) (*pb.Profession, error)
	Update(ctx context.Context, entity *pb.Profession) (id string, err error)
	Delete(ctx context.Context, entity *pb.Id) (id string, err error)
}

type AttributeI interface {
	CreateAttribute(ctx context.Context, entity *pb.CreateAttributeRequest) (id string, err error)
	GetAllAttribute(ctx context.Context, req *pb.GetAllAttributeRequest) (*pb.GetAllAttributeResponse, error)
	GetByIdAttribute(ctx context.Context, req *pb.AttributeId) (*pb.Attribute, error)
	UpdateAttribute(ctx context.Context, entity *pb.Attribute) (id string, err error)
	DeleteAttribute(ctx context.Context, entity *pb.AttributeId) (id string, err error)
}

type CompanyI interface {
	Create(ctx context.Context, entity *pc.CreateCompanyRequest) (id string, err error)
	GetAll(ctx context.Context, req *pc.GetAllCompanyRequest) (*pc.GetAllCompanyResponse, error)
	GetById(ctx context.Context, req *pc.Id) (*pc.Company, error)
	Update(ctx context.Context, entity *pc.Company) (id string, err error)
	Delete(ctx context.Context, entity *pc.Id) (id string, err error)
}

type PositionI interface {
	Create(ctx context.Context, entity *pb.CreatePositionRequest) (*pb.Position, error)
	GetAll(ctx context.Context, req *pb.GetAllPositionRequest) (*pb.GetAllPositionResponse, error)
	GetById(ctx context.Context, req *pb.PositionId) (*pb.Position, error)
	Update(ctx context.Context, entity *pb.UpdatePositionRequest) (*pb.PositionId, error)
	Delete(ctx context.Context, entity *pb.PositionId) (*pb.PositionId, error)
}
