package main

/********************
TODO 2024.7.4:
    [ ] Functional (SERVER)
        [x] Log Runs of a Task
        [x] Store in DB (SQLite3)
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
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

const base = "http://"
const url  = "localhost:"
const port = "8080"

var db, dberr = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})

type TaskLogLine struct {
    Message string
    TaskRun TaskRun `gorm:"serializer:json"`
}

type TaskRun struct {
    gorm.Model
    ID uint
    Method string
    TaskName string
    IncomingPayload any `gorm:"serializer:json"`
    Status string
    ErrorDetails error `gorm:"serializer:json"`
    TaskLogLines []TaskLogLine `gorm:"serializer:json"`
    Start time.Time
    Stop time.Time
}


func (task *TaskRun) log(message string) {
    logLine := TaskLogLine{}
    logLine.Message = message

    task.TaskLogLines = append(task.TaskLogLines, logLine)
    log.Println(task.TaskLogLines)
    log.Println(task.TaskName, "\n\t" + message)
}

func main() {
    if dberr != nil {
        panic("Connection Error: Failed to connect database")
    }
    db.AutoMigrate(&TaskRun{})

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
    run.TaskName = taskId
    run.IncomingPayload = payload
    run.Method = request.Method
    run.Status = "Pending..."
    log.Println(taskRunner(run))
}

func handlePostRequest(response http.ResponseWriter, request *http.Request) {
    // TODO: determine and type post data and delegate accordingly
}

func taskRunner(run TaskRun) TaskRun {
    start := time.Now()
    var err error = nil
    if run.TaskName == "exampleTaskA" {
        run, err = exampleTaskA(run)
    } else if run.TaskName == "exampleTaskB" {
        run, err = exampleTaskB(run)
    } else {
        run.TaskName = run.TaskName + ": Not found"
        err = errors.New("Task id provided is not registered in the system")
        run.ErrorDetails = err
    }
    if err != nil {
        run.Status = "Failed"
    } else {
        run.Status = "Success"
    }
    run.ErrorDetails = err
    run.Start = start
    run.Stop = time.Now()
    db.Create(&run)
    return run
}

func routineRunner() {
    go exampleRoutineA()
    go exampleRoutineB()
}

