package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func CreateBranchHandler(w http.ResponseWriter, r *http.Request) {

	varMap := mux.Vars(r)
	slugUrl, found := varMap["slug"]
	if !found {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println("not found")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	var thread models.Thread

	err = json.Unmarshal(body, &thread)
	if err != nil {
		if strings.HasPrefix(err.Error(), `parsing time "{}"`){
			thread.Created = time.Time{}
		} else{
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(err.Error())
			return
		}
	}

	var slugIsEmpty bool
	if len(thread.Slug) == 0 {
		slugIsEmpty = true
		thread.Slug = strings.Replace(strings.ToLower(thread.Title), " ", "_", -1)
	}

	thread.Forum = slugUrl

	t, err := models.GetInstance().CreateThread(thread)
	if err != nil {
		if err.Error() == errorPqNoDataFound {
			myJSON := fmt.Sprintf(`{"%s%s%s or %s%s"}`,
				messageCantFind, cantFindUser, thread.Author, cantFindForum, thread.Forum)
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

	if slugIsEmpty{
		t.Slug = ""
	}

	data, err := json.Marshal(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	if t.IsNew == false {
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
