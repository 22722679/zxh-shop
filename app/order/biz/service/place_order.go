package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/order/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/order/biz/model"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	if len(req.Items) == 0 {
		err = kerrors.NewBizStatusError(500001, "items is empty")
		return
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewUUID()
		o := &model.Order{
			OrderId:      orderId.String(),
			UserId:       uint(req.UserId),
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.StreetAddress = a.StreetAddress
			o.Consignee.City = a.City
			o.Consignee.Country = a.Country
		}
		err := tx.Create(o).Error
		if err != nil {
			return err
		}
		var items []model.OrderItem
		for _, v := range req.Items {
			items = append(items, model.OrderItem{
				OrderIdRefer: orderId.String(),
				ProductId:    v.Item.ProductId,
				Quantity:     int32(v.Item.Quantity),
				Cost:         v.Cost,
			})
		}

		if err := tx.Create(items).Error; err != nil {
			return err
		}
		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}
		return nil
	})
	return
}
