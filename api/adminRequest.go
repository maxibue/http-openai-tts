package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/maximierung/http-openai-tts/database"
	"github.com/maximierung/http-openai-tts/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func AdminRequest(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	config, _ := utils.LoadConfig()
	key := r.Header.Get("KEY")

	if key != config.AdminKey {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("{\"status\": \"Key error\",\"message\": \"The key provided in the 'KEY' header doesn't match the admin key.\"}"))
		return
	}

	requestType := strings.ToLower(r.PathValue("type"))

	if !utils.CheckAdminType(requestType) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"status\": \"Input error\",\"message\": \"The selected path doesn't exist.\"}"))
		return
	}
	key_name := r.FormValue("name")
	new_key := r.FormValue("key")

	switch requestType {
	case "add":
		if key_name == "" || key == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{\"status\": \"Input error\",\"message\": \"The 'key' and 'name' values can't be empty.\"}"))
			return
		}
		if database.AddKey(client, config.DBName, key_name, new_key) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("{\"status\": \"OK\",\"message\": \"Successfully created the key: %s\", \"key\": \"%s\"}", key_name, new_key)))
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"status\": \"Internal error\",\"message\": \"An error occurred while creating the key.\"}"))
			return
		}
	case "remove":
		if key_name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{\"status\": \"Input error\",\"message\": \"The 'name' value can't be empty.\"}"))
			return
		}
		if database.RemoveKey(client, config.DBName, key_name) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("{\"status\": \"OK\",\"message\": \"Successfully removed the key: %s\"}", key_name)))
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"status\": \"Internal error\",\"message\": \"An error occurred while removing the key.\"}"))
			return
		}
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("{\"status\": \"Internal error\",\"message\": \"An unknown error occured.\"}"))
}
