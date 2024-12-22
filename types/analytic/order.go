package analytic

type OrderAnalytic struct {
	OrdersCount int     `json:"ordersCount"`
	TotalAmount float64 `json:"totalAmount"`
	StartTime   string  `json:"startTime"`
}
