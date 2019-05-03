package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"net/http"
)

func GetDataBaseInfoHandler(w http.ResponseWriter, r *http.Request) {

	database, err := models.GetInstance().GetDatabase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
	database.Post = database.Post - 1
	database.User = database.User - 1
	database.Forum = database.Forum - 1
	database.Thread = database.Thread - 1

	data, err := json.Marshal(database)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
}
