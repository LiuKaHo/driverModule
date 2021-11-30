// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/LiuKaHo/driverModule/conf"
	"github.com/LiuKaHo/driverModule/dao/data"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"product/internal/biz"
	"product/internal/server"
	"product/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	productRepo := data.NewProductRepo(dataData, logger)

	productUserCase := biz.NewProductUserCase(productRepo, logger)
	productService := service.NewProductService(productUserCase, logger)
	httpServer := server.NewHTTPServer(confServer, productService, logger)
	grpcServer := server.NewGRPCServer(confServer, productService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
