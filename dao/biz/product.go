package biz

import "context"

type Product struct {
	ID   int64
	Name string
}

type ProductRepo interface {
	ListProduct(ctx context.Context) ([]*Product, error)
}
