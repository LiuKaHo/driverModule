package data

import (
	"context"
	"github.com/LiuKaHo/driverModule/dao/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type ProductRepo struct {
	data *Data
	log  log.Logger
}

func NewProductRepo(data *Data, log log.Logger) biz.ProductRepo {
	return &ProductRepo{
		data: data,
		log:  log,
	}
}

func (productRepo *ProductRepo) ListProduct(ctx context.Context) ([]*biz.Product, error) {
	res, err := productRepo.data.db.Product.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]*biz.Product, 0)
	for _, r := range res {
		response = append(response, &biz.Product{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	return response, nil
}
