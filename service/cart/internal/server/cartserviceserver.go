// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: cart.proto

package server

import (
	"context"

	"mall/service/cart/internal/logic"
	"mall/service/cart/internal/svc"
	"mall/service/cart/proto/cart"
)

type CartServiceServer struct {
	svcCtx *svc.ServiceContext
	cart.UnimplementedCartServiceServer
}

func NewCartServiceServer(svcCtx *svc.ServiceContext) *CartServiceServer {
	return &CartServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *CartServiceServer) AddItem(ctx context.Context, in *cart.AddItemReq) (*cart.AddItemResp, error) {
	l := logic.NewAddItemLogic(ctx, s.svcCtx)
	return l.AddItem(in)
}

func (s *CartServiceServer) GetCart(ctx context.Context, in *cart.GetCartReq) (*cart.GetCartResp, error) {
	l := logic.NewGetCartLogic(ctx, s.svcCtx)
	return l.GetCart(in)
}

func (s *CartServiceServer) EmptyCart(ctx context.Context, in *cart.EmptyCartReq) (*cart.EmptyCartResp, error) {
	l := logic.NewEmptyCartLogic(ctx, s.svcCtx)
	return l.EmptyCart(in)
}
