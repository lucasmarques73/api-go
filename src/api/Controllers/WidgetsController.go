package Controllers

import (
	"api/Errors"
	"api/Models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAllWidgets -  Listing all widgets
func GetAllWidgets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var widgets []Models.Widget
	if err := Models.WidgetsModel.Find().All(&widgets); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Errors:  true,
			Data:    "",
			Message: "Wigets not found",
		})
		return
	}

	json.NewEncoder(w).Encode(Response{
		Errors:  false,
		Data:    widgets,
		Message: "List of all wigets",
	})

}

// GetWidget - Listing a widget
func GetWidget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	idS := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idS)

	var widget Models.Widget
	res := Models.WidgetsModel.Find(id)
	err := res.One(&widget)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Errors:  true,
			Data:    "",
			Message: "Widget not found",
		})
		return
	}

	json.NewEncoder(w).Encode(Response{
		Errors:  false,
		Data:    widget,
		Message: "Widget data of id " + idS,
	})

}

// CreateWidget - Creating a widget
func CreateWidget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var widget Models.Widget
	_ = json.NewDecoder(r.Body).Decode(&widget)

	n := Models.WidgetsModel.Find("name", widget.Name)

	if count, _ := n.Count(); count > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Errors:  true,
			Data:    "",
			Message: "The name field must be unique",
		})
		return
	}

	res, err := Models.WidgetsModel.Insert(widget)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		Errors.CheckErr(err)
		return
	}

	widget.ID = res.(int64)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Errors:  false,
		Data:    widget,
		Message: "Widget created",
	})

}

// UpdateWidget - Updating a widget
func UpdateWidget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var widget Models.Widget
	res := Models.WidgetsModel.Find(id)
	err := res.One(&widget)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Errors:  true,
			Data:    "",
			Message: "Widget not found",
		})
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&widget)

	// Validate duplicate name
	// nid is Name and ID
	n := Models.WidgetsModel.Find("name", widget.Name)
	count, _ := n.Count()
	var nid Models.Widget
	n.One(&nid)
	if count > 0 && widget.ID != nid.ID {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Errors:  true,
			Data:    "",
			Message: "The name field must be unique",
		})
		return
	}

	if err = res.Update(widget); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		Errors.CheckErr(err)
		return
	}

	json.NewEncoder(w).Encode(Response{
		Errors:  false,
		Data:    widget,
		Message: "User updated",
	})

}

// DeleteWidget - Deleting a widget
func DeleteWidget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var widget Models.Widget
	res := Models.WidgetsModel.Find(id)
	err := res.One(&widget)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Errors:  false,
			Data:    "",
			Message: "Widget not found",
		})
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&widget)

	if err = res.Delete(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		Errors.CheckErr(err)
		return
	}

	json.NewEncoder(w).Encode(Response{
		Errors:  false,
		Data:    widget,
		Message: "Widget deleted",
	})

}
