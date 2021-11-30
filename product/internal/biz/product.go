package biz

import (
	"context"
	"github.com/LiuKaHo/driverModule/dao/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type ProductUserCase struct {
	repo biz.ProductRepo
	log  *log.Helper
}

func NewProductUserCase(repo biz.ProductRepo, logger log.Logger) *ProductUserCase {
	return &ProductUserCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ProductUserCase) ListProduct(ctx context.Context) (product []*biz.Product, err error) {
	return uc.repo.ListProduct(ctx)
}
