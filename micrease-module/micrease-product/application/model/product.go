package model

//模型定义
type Product struct {
	ID   uint
	Name string
}

func (this Product) TableName() string {
	return "tb_product"
}

func (this Product) GetID() uint {
	return this.ID
}
