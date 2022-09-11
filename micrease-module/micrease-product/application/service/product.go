package service

import (
	"errors"
	"fmt"
	pb "github.com/micrease/meshop-protos/pb/product"
	"github.com/micrease/micrease-core/context"
	"github.com/micrease/micrease-core/errs"
	"meshop-product-service/application/model"
	"meshop-product-service/application/repo"
)

/**
 * 业务逻辑
 */
type Product struct {
	repo *repo.Product
}

func NewProduct() *Product {
	product := &Product{}
	product.repo = repo.NewProduct()
	return product
}

func (this *Product) PageList(ctx *context.Context, req *pb.ProductPageReq, resp *pb.ProductPageResp) error {
	result, err := this.repo.Paginate(int(req.PageNo), int(req.PageSize))
	errs.PanicIf(err, 5001, "查询商品列表失败")
	fmt.Println(result)
	return nil
}

func (this *Product) Create(ctx *context.Context, req *pb.ProductInsertReq, resp *pb.ProductResp) error {
	newProd := model.Product{Name: req.Name}
	err := this.repo.Create(&newProd).Error
	errs.PanicIf(err, 5003, "创建商品失败")
	return nil
}

func (this *Product) Update(ctx *context.Context, req *pb.ProductUpdateReq, resp *pb.ProductResp) error {
	prod := &model.Product{ID: uint(req.Id), Name: req.Name}
	err := this.repo.Updates(prod).Error
	errs.PanicIf(err, 5004, "更新商品失败")
	resp.Data = &pb.Product{Id: int32(prod.ID), Name: prod.Name}
	return nil
}

func (this *Product) Delete(ctx *context.Context, req *pb.ProductDeleteReq, resp *pb.ProductResp) error {
	delProd := &model.Product{ID: uint(req.Id)}
	err := this.repo.Delete(delProd).Error
	errs.PanicIf(err, 5005, "删除商品失败")
	resp.Data = &pb.Product{Id: int32(delProd.ID), Name: delProd.Name}
	return nil
}

func (this *Product) Detail(ctx *context.Context, req *pb.ProductDetailReq, resp *pb.ProductResp) error {
	if prod, err := this.repo.Where("id=?", req.Id).GetOne(); err != nil {
		//resp.Data = &pb.Product{Id: 1, Name: "xxxxx"}
		return errors.New("error xxxx")
	} else {
		resp.Data = &pb.Product{Id: int32(prod.ID), Name: prod.Name}
	}
	return nil
}
