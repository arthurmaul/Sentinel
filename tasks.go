package main 

func exampleTaskA(run TaskRun) (TaskRun, error) {
    run.log("taska log 1...")
    run.log("taska log 2...")
    return run, nil
}

func exampleTaskB(run TaskRun) (TaskRun, error) {
    run.log("taskb log 1...")
    run.log("taskb log 2...")
    return run, nil
}

