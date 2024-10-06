package entity

import "github.com/mercadopago/sdk-go/pkg/payment"

type OrderState string
type Order struct {
	ID          string         `json:"id" db:"id"`
	Status      string         `json:"status" db:"status"`
	TotalAmount int            `json:"totalAmount" db:"total_amount"`
	Table       string         `json:"table" db:"board"`
	Items       MenuItemsSlice `json:"items,omitempty" db:"items"`
	PaymentID   string         `json:"paymentId" db:"payment_id"`
}

type OrderItem struct {
	ID       string `json:"id" db:"id"`
	OrderID  string `json:"orderId" db:"order_id"`
	ItemID   string `json:"itemId" db:"item_id"`
	Comments string `json:"comments" db:"comments"`
	Quantity int    `json:"quantity" db:"quantity"`
	Price    int    `json:"price" db:"price"`
}

type CreateOrderRequest struct {
	ID      string          `json:"id"`
	Payment payment.Request `json:"payment"`
	Table   string          `json:"table"`
	Items   []OrderItem     `json:"items"`
}

type UpdateStatusRequest struct {
	Status OrderState `json:"status"`
}
