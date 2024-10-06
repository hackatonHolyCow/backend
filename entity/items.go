package entity

import (
	"encoding/json"

	"github.com/lib/pq"
)

type MenuItem struct {
	ID          string         `json:"id" db:"id"`
	Name        string         `json:"name" db:"name"`
	Description string         `json:"description" db:"description"`
	Price       int            `json:"price" db:"price"`
	Tags        pq.StringArray `json:"tags" db:"tags"`
	Picture     string         `json:"picture" db:"picture"`
	Comments    string         `json:"comments,omitempty" db:"comments"`
	Quantity    int            `json:"quantity,omitempty" db:"quantity"`
}

type MenuItemsSlice []*MenuItem

func (m *MenuItemsSlice) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	if err := json.Unmarshal(src.([]byte), m); err != nil {
		return err
	}

	return nil
}
