package widgets

import (
	"api/errors"
	"database/sql"

	_ "github.com/lib/pq"
)

type Widget struct {
	ID        int     `db:"id" json:"id,omitempty"`
	Name      string  `db:"name" json:"name,omitempty"`
	Color     string  `db:"color" json:"color,omitempty"`
	Price     float64 `db:"price" json:"price,omitempty"`
	Melts     bool    `db:"melts" json:"melts,omitempty"`
	Inventory int     `db:"inventory" json:"inventory,omitempty"`
}

var conStr = "host=db port=5432 user=go password=go dbname=go sslmode=disable"
var db, err = sql.Open("postgres", conStr)

func ListWidgets() []Widget {

	rows, err := db.Query("SELECT * FROM widgets")
	errors.CheckErr(err)

	widgets := []Widget{}

	for rows.Next() {
		widget := Widget{}
		rows.Scan(&widget.ID, &widget.Name, &widget.Color, &widget.Price, &widget.Melts, &widget.Inventory)
		widgets = append(widgets, widget)
	}

	return widgets
}
func GetWidgetById(id string) Widget {

	row := db.QueryRow("SELECT * FROM widgets WHERE id =" + id)
	widget := Widget{}
	row.Scan(&widget.ID, &widget.Name, &widget.Color, &widget.Price, &widget.Melts, &widget.Inventory)

	return widget
}
func CreateWidgets(widget Widget) Widget {
	return widget
}
func UpdateWidgets(widget Widget) Widget {
	return widget
}
func DeleteWidgets() {}
