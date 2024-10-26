package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	// GetTotal() (int, error)
	FindOrderById(id string) (*Order, error)
	FindOrders() ([]*Order, error)
}
