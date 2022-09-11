package handler

import (
	"context"
	pb "github.com/micrease/meshop-protos/pb/product"
	"go-micro.dev/v4"
	"meshop-product-service/application/service"
)

type Product struct {
	//初始化上下文传递的通用变量，比如log,gorm,redis等
	RpcHandler
	//初始化
	service *service.Product
}

//把自身注册到rpc handler中
func RegisterProduct(svr micro.Service) {
	s := &Product{}
	//1,创建上下文
	s.NewContext()
	s.service = service.NewProduct()
	//2,把当前对象注册到rpc中
	pb.RegisterProductServiceHandler(svr.Server(), s)
}

func (this *Product) Create(ctx context.Context, req *pb.ProductInsertReq, resp *pb.ProductResp) (err error) {
	return this.service.Create(this.ctx, req, resp)
}

func (this *Product) Update(ctx context.Context, req *pb.ProductUpdateReq, resp *pb.ProductResp) (err error) {
	return this.service.Update(this.ctx, req, resp)
}

func (this *Product) Delete(ctx context.Context, req *pb.ProductDeleteReq, resp *pb.ProductResp) (err error) {
	return this.service.Delete(this.ctx, req, resp)
}

func (this *Product) Detail(ctx context.Context, req *pb.ProductDetailReq, resp *pb.ProductResp) (err error) {
	return this.service.Detail(this.ctx, req, resp)
}

func (this *Product) PageList(ctx context.Context, req *pb.ProductPageReq, resp *pb.ProductPageResp) (err error) {
	return this.service.PageList(this.ctx, req, resp)
}
