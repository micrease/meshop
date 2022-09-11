package handler

import (
	"context"
	me "github.com/micrease/micrease-core/context"
	"meshop-product-service/datasource"
)

type Handler interface {
	NewContext()
	NewWithContext(ctx context.Context)
}

type RpcHandler struct {
	Handler
	ctx *me.Context
}

func (this *RpcHandler) NewContext() {
	this.ctx = new(me.Context)
	this.ctx.Orm = datasource.GetDB()
}

func (this *RpcHandler) NewWithContext(ctx context.Context) {
	if this.ctx == nil {
		this.ctx = new(me.Context)
		this.ctx.Orm = datasource.GetDB()
	}
	this.ctx.Ctx = ctx
}
