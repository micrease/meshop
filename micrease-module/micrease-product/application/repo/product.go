package repo

import (
	"github.com/micrease/gorme"
	"meshop-product-service/application/model"
	"meshop-product-service/datasource"
)

/**
 * 数据库操作
 */
type Product struct {
	gorme.Repository[model.Product]
}

func NewProduct() *Product {
	product := new(Product)
	product.SetDB(datasource.GetDB())
	return product
}

//todo::在下面扩展你的自定义方法吧
func (this *Product) FindByName(name string) ([]model.Product, error) {
	return this.Where("name like ?", "%"+name+"%").List()
}
