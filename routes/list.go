package routes

import (
	"encoding/json"
	"net/http"

	"bey.com/models"
)

// ListHandler list page handler
func ListHandler(rw http.ResponseWriter, rq *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	data := make(map[string]string, 0)
	data["id"] = "ID00001"
	data["name"] = "bey"
	data["gender"] = "male"

	res := models.Response{
		Code: 200,
		Msg:  "success!",
		Data: data,
	}

	json, err := json.Marshal(res)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Write(json)
}
