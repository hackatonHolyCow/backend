package entity

type OrderState string

const (
	OrderstateCancelled = "cancelled"
	OrderStateConfirmed = "confirmed"
	OrderStateFinished  = "finished"
)

type Order struct {
	ID          string     `json:"id" db:"id"`
	State       OrderState `json:"state" db:"state"`
	TotalAmount int        `json:"totalAmount" db:"total_amount"`
	Table       string     `json:"table" db:"table"`
}
type OrderItem struct {
	ID       string `json:"id" db:"id"`
	OrderID  string `json:"orderId" db:"order_id"`
	ItemID   string `json:"itemId" db:"item_id"`
	Comments string `json:"comments" db:"comments"`
	Quantity string `json:"quantity" db:"quantity"`
}
