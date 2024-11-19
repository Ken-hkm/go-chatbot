package tools

import (
	_ "github.com/tmc/langchaingo/chains"
	_ "github.com/tmc/langchaingo/llms"
	_ "github.com/tmc/langchaingo/tools"
	"gorm.io/gorm"
)

type OrderTool struct {
	db *gorm.DB
}

func NewOrderTool(db *gorm.DB) *OrderTool {
	return &OrderTool{
		db: db,
	}
}

func (o *OrderTool) Name() string {
	return "Order Tool"
}

func (o *OrderTool) Description() string {
	return "Gets and fetch all Order Information by user Id"
}

//
//func (o *OrderTool) Run(ctx context.Context, query string) (string, error) {
//	//orderRepo := repository.NewOrderRepository(o.db)
//	//orderService := service.NewOrderService(orderRepo)
//	//orderData, err := orderService.GetUserByID(1)
//	//if err != nil {
//	//	return "", err
//	//}
//	//
//	//return orderData[0].SalesNo, nil
//}
