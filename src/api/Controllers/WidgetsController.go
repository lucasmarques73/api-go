package controllers

import (
	"api/Errors"
	"api/Models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllWidgets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var widgets []models.Widget
	if err := models.WidgetsModel.Find().All(&widgets); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"errors":  strconv.FormatBool(true),
			"data":    "",
			"message": "Wigets not found",
		})
	} else {
		data, _ := json.Marshal(widgets)
		json.NewEncoder(w).Encode(map[string]string{
			"errors":  strconv.FormatBool(false),
			"data":    string(data),
			"message": "List of all wigets",
		})
	}
}

func GetWidget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	idS := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idS)

	var widget models.Widget
	res := models.WidgetsModel.Find(id)
	err := res.One(&widget)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   strconv.FormatBool(true),
			"data":    "",
			"message": "Widget not found",
		})
	} else {
		data, _ := json.Marshal(widget)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   strconv.FormatBool(false),
			"data":    string(data),
			"message": "Widget data of id " + idS,
		})
	}
}

func CreateWidget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var widget models.Widget
	_ = json.NewDecoder(r.Body).Decode(&widget)

	n := models.WidgetsModel.Find("name", widget.Name)

	if count, _ := n.Count(); count > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   strconv.FormatBool(true),
			"data":    "",
			"message": "The name field must be unique",
		})
	} else {
		if res, err := models.WidgetsModel.Insert(widget); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errors.CheckErr(err)
		} else {
			widget.ID = res.(int64)
			w.WriteHeader(http.StatusCreated)
			data, _ := json.Marshal(widget)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   strconv.FormatBool(false),
				"data":    string(data),
				"message": "Widget created",
			})
		}
	}

}

func UpdateWidget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var widget models.Widget
	res := models.WidgetsModel.Find(id)
	err := res.One(&widget)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   strconv.FormatBool(true),
			"data":    "",
			"message": "Widget not found",
		})
	} else {
		_ = json.NewDecoder(r.Body).Decode(&widget)

		// Validate duplicate name
		// nid is Name and ID
		n := models.WidgetsModel.Find("name", widget.Name)
		count, _ := n.Count()
		var nid models.Widget
		n.One(&nid)
		if count > 0 && widget.ID != nid.ID {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   strconv.FormatBool(true),
				"data":    "",
				"message": "The name field must be unique",
			})
		} else {
			if err = res.Update(widget); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				errors.CheckErr(err)
			} else {
				data, _ := json.Marshal(widget)
				json.NewEncoder(w).Encode(map[string]string{
					"error":   strconv.FormatBool(false),
					"data":    string(data),
					"message": "User updated",
				})
			}
		}
	}
}

func DeleteWidget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var widget models.Widget
	res := models.WidgetsModel.Find(id)
	err := res.One(&widget)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   strconv.FormatBool(false),
			"data":    "",
			"message": "Widget not found",
		})
	} else {
		_ = json.NewDecoder(r.Body).Decode(&widget)

		if err = res.Delete(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errors.CheckErr(err)
		} else {
			data, _ := json.Marshal(widget)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   strconv.FormatBool(false),
				"data":    string(data),
				"message": "Widget deleted",
			})
		}
	}
}
