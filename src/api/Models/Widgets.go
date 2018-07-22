package models

import "api/Db"

// Widget struct of table widgets
type Widget struct {
	ID        int64   `db:"id,omitempty" json:"id,omitempty"`
	Name      string  `validate:"required" db:"name" json:"name,omitempty"`
	Color     string  `db:"color" json:"color,omitempty"`
	Price     float64 `db:"price" json:"price,omitempty"`
	Melts     bool    `db:"melts" json:"melts,omitempty"`
	Inventory int     `validate:"required" db:"inventory" json:"inventory,omitempty"`
}

// WidgetsModel Ppinting to the "widgets" table
var WidgetsModel = db.Sess.Collection("widgets")
