package main

import (
	"net/http"
	"time"
)

func exampleRoutineA() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		<-ticker.C
		http.Get(base + url + port + "/exampleTaskA?hello=world_a")
	}
}

func exampleRoutineB() {
	ticker := time.NewTicker(1 * time.Minute)
	for {
		<-ticker.C
		http.Get(base + url + port + "/exampleTaskB?testing=routine_b")
	}
}
