package routes

import (
    "net/http"
    "github.com/user/cv/handlers"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        handlers.Index,
    },
    Route{
        "Candidate",
        "GET",
        "/candidate/{id}",
        handlers.GetCandidate,
    },
    Route{
        "CreateCandidate",
        "POST",
        "/candidate",
        handlers.CreateCandidate,
    },
    Route{
        "DeleteCandidate",
        "DELETE",
        "/candidate/{id}",
        handlers.DeleteCandidate,
    },
}
