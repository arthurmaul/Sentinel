package main

import (
	"net/http"
	"time"
)

func fireRoutine(interval time.Duration, route string) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		http.Get(base + url + port + route)
	}
}

func exampleRoutineA() {
    fireRoutine(5 * time.Second, "/exampleTaskA?hello=world_a")
}

func exampleRoutineB() {
    fireRoutine(1 * time.Minute, "/exampleTaskB?testing=routine_b")
}
