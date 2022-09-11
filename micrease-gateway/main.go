package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/micrease/meshop/micrease-protos/gw/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

var (
	productEndpoint = flag.String("product_endpoint", "localhost:8004", "product service地址")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	serveMux := runtime.NewServeMux()

	err := product.RegisterProductServiceHandlerFromEndpoint(ctx, serveMux, *productEndpoint, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})

	if err != nil {
		return err
	}
	return http.ListenAndServe(":8080", serveMux)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
