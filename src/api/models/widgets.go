package models

import (
	// driver for database
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

// func ListWidgets() []Widget {
// 	var db, err = sql.Open("postgres", conStr)
// 	rows, err := db.Query("SELECT * FROM widgets")
// 	errors.CheckErr(err)

// 	widgets := []Widget{}

// 	for rows.Next() {
// 		widget := Widget{}
// 		rows.Scan(&widget.ID, &widget.Name, &widget.Color, &widget.Price, &widget.Melts, &widget.Inventory)
// 		widgets = append(widgets, widget)
// 	}
// 	db.Close()

// 	return widgets
// }
// func GetWidgetById(id string) Widget {
// 	var db, err = sql.Open("postgres", conStr)
// 	row := db.QueryRow("SELECT * FROM widgets WHERE id =" + id)
// 	widget := Widget{}
// 	row.Scan(&widget.ID, &widget.Name, &widget.Color, &widget.Price, &widget.Melts, &widget.Inventory)
// 	db.Close()

// 	return widget
// }
// func CreateWidgets(widget Widget) Widget {
// 	var db, _ = sql.Open("postgres", conStr)
// 	_, err := db.Exec("INSERT INTO widgets (name,color,price,melts,inventory) VALUES ('" + widget.Name + "','" + widget.Color + "'," + strconv.FormatFloat(widget.Price, 'E', 2, 64) + "," + strconv.FormatBool(widget.Melts) + "," + strconv.Itoa(widget.Inventory) + ")")
// 	errors.CheckErr(err)
// 	db.Close()
// 	return widget
// }
// func UpdateWidgets(widget Widget, id string) Widget {
// 	var db, _ = sql.Open("postgres", conStr)
// 	_, err := db.Exec("UPDATE widgets SET name=" + widget.Name + ",color=" + widget.Color + ",price=" + strconv.FormatFloat(widget.Price, 'E', 2, 64) + ",melts=" + strconv.FormatBool(widget.Melts) + ",inventory=" + strconv.Itoa(widget.Inventory) + " WHERE id=" + id)
// 	errors.CheckErr(err)
// 	db.Close()
// 	return widget
// }
// func DeleteWidgets(id string) {
// 	var db, _ = sql.Open("postgres", conStr)
// 	_, err := db.Exec("DELETE FROM widgets WHERE id = " + id)
// 	errors.CheckErr(err)
// 	db.Close()
// }
