package main

import (
    "net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)

func GetOnly(h handler) handler {

    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            h(w, r)
            return
        }
        http.Error(w, "get only", http.StatusMethodNotAllowed)
    }
}

func PostOnly(h handler) handler {

    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            h(w, r)
            return
        }
        http.Error(w, "post only", http.StatusMethodNotAllowed)
    }
}
