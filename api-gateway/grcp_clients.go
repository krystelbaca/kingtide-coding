package main

import (
	"log"

	pbOrder "github.com/krystelbaca/kingtide-coding/api-gateway/proto/order"
	pbProduct "github.com/krystelbaca/kingtide-coding/api-gateway/proto/product"
	"google.golang.org/grpc"
)

// GRPCClients holds the clients for each gRPC service
type GRPCClients struct {
	ProductClient pbProduct.ProductServiceClient
	OrderClient   pbOrder.OrderServiceClient
}

// NewGRPCClients creates new gRPC clients for product and order services
func NewGRPCClients(productAddr, orderAddr string) *GRPCClients {
	// Set up a connection to the product service.
	productConn, err := grpc.Dial(productAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	productClient := pbProduct.NewProductServiceClient(productConn)

	// Set up a connection to the order service.
	orderConn, err := grpc.Dial(orderAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	orderClient := pbOrder.NewOrderServiceClient(orderConn)

	return &GRPCClients{
		ProductClient: productClient,
		OrderClient:   orderClient,
	}
}
