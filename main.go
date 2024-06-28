package main

import (
    "log"
    "net/http"
)

func handleTask(response http.ResponseWriter, request *http.Request) {
    method  := request.Method
    taskID  := request.PathValue("taskid")
    payload := make(map[string][]string)

    // Determine method type and extract data.
    switch method {
    case "GET":
        payload = request.URL.Query()
    }

    // Call proper function with payload data.
    switch taskID {
    case "polkadots":
        polkadots(payload)
    }

    log.Println(method, "method targeting task", taskID)
    log.Println("payload of ", payload)
}

func polkadots(data any){
    log.Println("!!!!!!!")
    log.Println(data)
}

func main() {
    router := http.NewServeMux()
    router.HandleFunc("GET /tasks/{taskid}", handleTask)
    router.HandleFunc("POST /tasks/{taskid}", handleTask)
    server := http.Server{
        Addr: "localhost:8080",
        Handler: router,
    }

    log.Println("Server Created")
    log.Println("Running on Port: 8080")
    server.ListenAndServe()
}
