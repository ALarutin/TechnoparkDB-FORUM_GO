package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ChangeMessageHandler(w http.ResponseWriter, r *http.Request) {

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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	var post models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	post, err = models.GetInstance().UpdatePost(post.Message, id)
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

	data, err := json.Marshal(post)
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
