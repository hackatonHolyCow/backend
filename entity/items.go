package entity

import "github.com/lib/pq"

type MenuItem struct {
	ID          string         `json:"id" db:"id"`
	Name        string         `json:"name" db:"name"`
	Description string         `json:"description" db:"description"`
	Price       float64        `json:"price" db:"price"`
	Tags        pq.StringArray `json:"tags" db:"tags"`
	Picture     string         `json:"picture" db:"picture"`
}
