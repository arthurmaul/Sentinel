package main

/********************
TODO 2024.7.4:
    [ ] Functional (SERVER)
        [x] Log Runs of a Task
        [ ] Store in DB (SQLite3)
        [ ] API/View URL to pull possible tasks (think like http://ourgoserver.com/tasks) and render HTMX layout
        [ ] API/View URL to show a previous task run and it's details

    [ ] UI (HTMX/SERVER):
        [ ] Display list of Tasks that can be run
        [ ] Next to each task display, show historical runs
        [ ] Click a task name to view list of past runs
        [ ] Past run should log STDOUT, STDERR, incoming payload, etc, Start, End

********************/

import (
    "log"
    "errors"
    "net/http"
    "time"
)

const base = "http://"
const url  = "localhost:"
const port = "8080"

type TaskLogLine struct {
    time time.Time
    message string
}

type TaskRun struct {
    method string
    task string
    incomingPayload any
    status string
    errorDetails error
    logs []TaskLogLine
    start time.Time
    stop time.Time
}

func (task *TaskRun) log(message string) {
    logLine := TaskLogLine{}
    logLine.message = message
    logLine.time = time.Now()

    task.logs = append(task.logs, logLine)
    log.Println(task.logs)
    log.Println(task.task, "\n\t" + message)
}

func (t *TaskRun) save() {
    //Gorm save t.save()
}

func main() {
    router := http.NewServeMux()
    router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
        http.ServeFile(response, request, "static/index.html")
    })
    router.HandleFunc("GET /{taskid}", handleGetRequest)
    router.HandleFunc("POST /{taskid}", handlePostRequest)
    server := http.Server{
        Addr: url + port,
        Handler: router,
    }

    routineRunner()
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
    run := TaskRun{}
    run.task = taskId
    run.incomingPayload = payload
    run.method = request.Method
    run.status = "Pending..."
    log.Println(taskRunner(run))
}

func handlePostRequest(response http.ResponseWriter, request *http.Request) {
    // TODO: determine and type post data and delegate accordingly
}

func taskRunner(run TaskRun) TaskRun {
    start := time.Now()
    var err error = nil
    if run.task == "exampleTaskA" {
        run, err = exampleTaskA(run)
    } else if run.task == "exampleTaskB" {
        run, err = exampleTaskB(run)
    } else {
        run.task = run.task + ": Not found"
        err = errors.New("Task id provided is not registered in the system")
        run.errorDetails = err
    }
    if err != nil {
        run.status = "Failed"
    } else {
        run.status = "Success"
    }
    run.errorDetails = err
    run.start = start
    run.stop = time.Now()
    // run.save()
    return run
}

func routineRunner() {
    go exampleRoutineA()
    go exampleRoutineB()
}

