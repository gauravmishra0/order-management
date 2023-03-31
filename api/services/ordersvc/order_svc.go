package ordersvc

var OrderSvc OrderActivities = NewOrderService()

type OrderActivities interface {
}

type OrderActivitiesStruct struct {
}

func NewOrderService() *OrderActivitiesStruct {
	order := new(OrderActivitiesStruct)
	return order
}
