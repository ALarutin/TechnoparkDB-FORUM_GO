package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func GetThreadInfoPostHandler(w http.ResponseWriter, r *http.Request) {

	varMap := mux.Vars(r)
	slug, found := varMap["id"]
	if !found {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println("not found")
		return
	}

	id, err := strconv.Atoi(slug)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	related := strings.Split(r.URL.Query().Get("related"), ",")

	postInfo, err := models.GetInstance().GetPostInfo(id, related)
	if err != nil {
		if err.Error() == errorPqNoDataFound {
			myJSON := fmt.Sprintf(`{"%s%s%v"}`, messageCantFind, cantFindPost, id)
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

	data, err := json.Marshal(postInfo)
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
