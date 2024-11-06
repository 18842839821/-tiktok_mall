// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: order.proto

package server

import (
	"context"

	"mall/service/order/internal/logic"
	"mall/service/order/internal/svc"
	"mall/service/order/proto/order"
)

type OrderServiceServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrderServiceServer
}

func NewOrderServiceServer(svcCtx *svc.ServiceContext) *OrderServiceServer {
	return &OrderServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServiceServer) PlaceOrder(ctx context.Context, in *order.PlaceOrderReq) (*order.PlaceOrderResp, error) {
	l := logic.NewPlaceOrderLogic(ctx, s.svcCtx)
	return l.PlaceOrder(in)
}

func (s *OrderServiceServer) ProcessOrder(ctx context.Context, in *order.ProcessOrderReq) (*order.ProcessOrderResp, error) {
	l := logic.NewProcessOrderLogic(ctx, s.svcCtx)
	return l.ProcessOrder(in)
}

func (s *OrderServiceServer) ListOrder(ctx context.Context, in *order.ListOrderReq) (*order.ListOrderResp, error) {
	l := logic.NewListOrderLogic(ctx, s.svcCtx)
	return l.ListOrder(in)
}

func (s *OrderServiceServer) MarkOrderPaid(ctx context.Context, in *order.MarkOrderPaidReq) (*order.MarkOrderPaidResp, error) {
	l := logic.NewMarkOrderPaidLogic(ctx, s.svcCtx)
	return l.MarkOrderPaid(in)
}
