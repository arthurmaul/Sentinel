package main

import (
    "time"
    "log"
    "errors"
    "net/http"
)

const base = "http://"
const url  = "localhost:"
const port = "8080"

func main() {
    router := http.NewServeMux()
    router.HandleFunc("GET /{taskid}", handleGetRequest)
    router.HandleFunc("POST /{taskid}", handlePostRequest)

    server := http.Server{
        Addr: url + port,
        Handler: router,
    }

    go exampleRoutineA()
    go exampleRoutineB()

    err := server.ListenAndServe()
    if err != nil {
        if !errors.Is(err, http.ErrServerClosed) {
            log.Printf("error running http server: %s\n", err)
        }
    }
}

func handleGetRequest(response http.ResponseWriter, request *http.Request) {
    taskId  := request.PathValue("taskid")
    payload := request.URL.Query()

    log.Println(
        "REQUEST",
        "\n\tmethod:",  request.Method,
        "\n\ttarget:",  taskId,
        "\n\tpayload:", payload,
    )

    triggerTask(taskId, payload)

}

func handlePostRequest(response http.ResponseWriter, request *http.Request) {
    // TODO: determine and type post data and delegate accordingly
}

func triggerTask(id string, payload any) {
    switch id {
    case "exampleTaskA":
        exampleTaskA(payload)
    case "exampleTaskB":
        exampleTaskB(payload)
    default:
        log.Println("Task not found.")
    }
}

func exampleTaskA(data any) error {
    log.Println("Task A running ...")
    return nil
}

func exampleTaskB(data any) error {
    log.Println("Task B running ...")
    return nil
}

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

