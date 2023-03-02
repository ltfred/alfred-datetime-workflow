package main

import (
	aw "github.com/deanishe/awgo"
	datetimeWorkflow "github.com/ltfred/alfred-datetime-workflow/datetime-workflow"
)

func main() {
	workflow := datetimeWorkflow.DatetimeWorkflow{Workflow: aw.New()}
	workflow.Run()
}
