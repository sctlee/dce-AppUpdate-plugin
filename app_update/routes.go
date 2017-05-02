package main

import (
    "net/http"

    "github.com/sctlee/app_update/handlers"
)

// Route ...
type Route struct {
    Name string
    Method string
    Pattern string
    HandlerFunc http.HandlerFunc
}

var routes = []Route{
    Route{
        "GetApp",
        "Get",
        "/api/app/{appName}",
        handlers.GetAppHandler,
    },
    Route{
        "UpdateApp",
        "PUT",
        "/api/app/{appName}",
        handlers.UpdateAppHandler,
    },
}
