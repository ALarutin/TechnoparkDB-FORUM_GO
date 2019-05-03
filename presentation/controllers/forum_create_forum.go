package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateForumHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	var forum models.Forum

	err = json.Unmarshal(body, &forum)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	f, err := models.GetInstance().CreateForum(forum)
	if err != nil {
		if err.Error() == errorPqNoDataFound {
			myJSON := fmt.Sprintf(`{"%s%s%s"}`, messageCantFind, cantFindUser, forum.User)
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte(myJSON))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Error.Println(err.Error())
				return
			}
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	data, err := json.Marshal(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	if f.IsNew == false {
		w.WriteHeader(http.StatusConflict)
		_, err = w.Write(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(err.Error())
			return
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
}
