package models

import "api/db"

// Widget struct of table widgets
type Widget struct {
	ID        int     `db:"id" json:"id,omitempty"`
	Name      string  `db:"name" json:"name,omitempty"`
	Color     string  `db:"color" json:"color,omitempty"`
	Price     float64 `db:"price" json:"price,omitempty"`
	Melts     bool    `db:"melts" json:"melts,omitempty"`
	Inventory int     `db:"inventory" json:"inventory,omitempty"`
}

// WidgetsModel Ppinting to the "widgets" table
var WidgetsModel = db.Sess.Collection("widgets")
