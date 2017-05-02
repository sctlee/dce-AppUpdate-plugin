package handlers

import (
    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux"

    "github.com/sctlee/app_update/update"
)

// GetAppHandler ...
func GetAppHandler(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    appName := vars["appName"]

    filteredServices, err := update.GetApp(appName)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(filteredServices); err != nil {
        panic(err)
    }
}

// UpdateAppHandler ...
func UpdateAppHandler(w http.ResponseWriter, r *http.Request)  {
    // vars := mux.Vars(r)
    // appName := vars["appName"]

    var jsonData []update.ServiceUpdateSpec
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    json.Unmarshal(body, &jsonData)

    // err := update.UpdateApp(appName, []update.ServiceUpdateSpec{})
    // if err != nil {
    //     w.WriteHeader(http.StatusBadRequest)
    //     return
    // }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(jsonData); err != nil {
        panic(err)
    }
    //
    // w.WriteHeader(http.StatusOK)
    // return
}
