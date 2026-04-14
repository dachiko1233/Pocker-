package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Player := strings.TrimPrefix(r.URL.Path, "/Player")
	fmt.Fprint(w, p.store.GetPlayerScore(Player))
}

func GetPlayerScore(name string) string {

	if name == "Peppr" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""

}
