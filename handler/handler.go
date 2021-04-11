package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"proxy/utils"
)

type Handler struct{}
type data map[string]interface{}

func JSONWriter(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (h *Handler) MicroserviceName(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		JSONWriter(w, data{
			"Error": "Method Not Allowed",
		}, http.StatusMethodNotAllowed)
		return
	}
	resp, err := http.Get("http://localhost:8082/microservice/name")
	if err != nil {
		JSONWriter(w, data{
			"Error": "Internal server error",
		}, http.StatusInternalServerError)
		return
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		JSONWriter(w, data{
			"Error": "Internal server error",
		}, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write(responseData)
}

func (h *Handler) User(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		JSONWriter(w, data{
			"Error": "Method Not Allowed",
		}, http.StatusMethodNotAllowed)
		return
	}
	userName := r.Header.Get("username")
	if userName == "" {
		JSONWriter(w, data{
			"Error": "no username found in header",
		}, http.StatusInternalServerError)
		return
	}
	resp, err := utils.HttpReqBuilder("http://localhost:8081/auth", userName)
	if err != nil {
		JSONWriter(w, data{
			"Error": "internal server error",
		}, http.StatusInternalServerError)
		return
	}
	if resp.Status == "200 OK" {
		userResp, err := utils.HttpReqBuilder("http://localhost:8082/user/profile", userName)
		if err != nil {
			JSONWriter(w, data{
				"Error": "internal server error",
			}, http.StatusInternalServerError)
			return
		}
		responseData, err := ioutil.ReadAll(userResp.Body)
		if err != nil {
			JSONWriter(w, data{
				"Error": "internal server error",
			}, http.StatusInternalServerError)
			return
		}
		if err != nil {
			JSONWriter(w, data{
				"Error": "Internal server error",
			}, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(200)
		w.Write(responseData)
		return

	} else {
		JSONWriter(w, data{"Response": "not authorized"}, 403)
		return
	}

}
