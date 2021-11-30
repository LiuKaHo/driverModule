package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "product/api/helloworld/v1"
	"product/internal/biz"
)

// GreeterService is a greeter service.
type ProductService struct {
	v1.UnimplementedProductServer

	uc  *biz.ProductUserCase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewProductService(uc *biz.ProductUserCase, logger log.Logger) *ProductService {
	return &ProductService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *ProductService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *ProductService) ListProduct(ctx context.Context, request *v1.ListProductRequest) (*v1.ListProductReply, error) {
	res, err := s.uc.ListProduct(ctx)
	if err != nil {
		return nil, err
	}
	response := &v1.ListProductReply{}

	for _, r := range res {
		response.Results = append(response.Results, &v1.Pro{
			Id:   r.ID,
			Name: r.Name,
		})
	}

	return response, nil
}
