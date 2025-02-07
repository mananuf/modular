package main

import (
	// "fmt"
	"fmt"

	"github.com/mananuf/task-manager/types"
)

func main() {
	dashboard := types.NewTask(
		"Create Dashboard",
		types.Frontend,
		types.High,
		types.Unassigned,
	)

	login := types.NewTask(
		"Create Login",
		types.Frontend,
		types.High,
		types.Unassigned,
	)

	userprofile := types.NewTask(
		"Create UserProfile",
		types.Frontend,
		types.Low,
		types.Unassigned,
	)

	database := types.NewTask(
		"Setup Database",
		types.Backend,
		types.High,
		types.Unassigned,
	)

	forgotPassword := types.NewTask(
		"reset password",
		types.Backend,
		types.Low,
		types.Unassigned,
	)
	
	tr := types.NewTaskRepository()

	tr.AddTaskToRepository(dashboard)
	tr.AddTaskToRepository(userprofile)
	tr.AddTaskToRepository(login)
	tr.AddTaskToRepository(database)
	tr.AddTaskToRepository(forgotPassword)
	tr.GetTaskByPriorityOrStatus(types.Frontend)
	types.PrintActualValues(tr.GetTaskByPriorityOrStatus(types.Backend))

	dashboard.UpdateTask(types.Assigned)

	fmt.Println(dashboard)
	types.PrintActualValues(tr.GetTaskByPriorityOrStatus(types.Assigned))
}