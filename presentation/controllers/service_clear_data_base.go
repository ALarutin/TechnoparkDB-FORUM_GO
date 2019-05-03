package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"net/http"
)

func ClearDataBaseHandler(w http.ResponseWriter, r *http.Request) {

	err := models.GetInstance().ClearDatabase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
}
