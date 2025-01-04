package backend

import "github.com/google/uuid"

type Order struct {
	OrderId  uuid.UUID
	Customer uuid.UUID
	Item     uuid.UUID
	Count    int
}

type SnackBackEnd struct {
	Orders []*Order
}

func NewSnackBackEnd() *SnackBackEnd {
	return &SnackBackEnd{
		Orders: make([]*Order, 0),
	}
}

func (s *SnackBackEnd) AddOrder(customerId uuid.UUID, itemId uuid.UUID, count int) uuid.UUID {
	orderID := uuid.New()
	s.Orders = append(s.Orders, &Order{
		OrderId:  orderID,
		Customer: customerId,
		Item:     itemId,
		Count:    count,
	})

	return orderID
}
