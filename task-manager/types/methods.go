package types

import (
	// "fmt"
)

func (tr *TaskRepository) GetTaskByPriorityOrStatus(value any) []*Task {
	var tasks []*Task

	for _, task := range tr.AllTasks {
		switch v := value.(type) {

		case TaskPriority:
			if v == task.Priority {
				tasks = append(tasks, task)
			}

		case TaskStatus:
			if v == task.Status {
				tasks = append(tasks, task)
			}

		case Teams:
			if v == task.TeamName {
				tasks = append(tasks, task)
			}

		default:
			panic("invalid query")
		}
	}

	return tasks
}

func (tr *TaskRepository) AddTaskToRepository(task *Task) {
	tr.AllTasks[uint(len(tr.AllTasks) + 1)] = task
}

func (t *Task) UpdateTask(value any) {
	switch v := value.(type) {
	case TaskPriority:
		t.Priority = v

	case TaskStatus:
		t.Status = v

	case Teams:
		t.TeamName = v

	default:
		panic("invalid type")
	}
}

// ------------------ CONSTRUCTOR METHODS ---------------------------
func NewTask(title string, teamName Teams, priority TaskPriority, status TaskStatus) *Task {
	return &Task{
		Title:    title,
		TeamName: teamName,
		Priority: priority,
		Status:   status,
	}
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		AllTasks: make(Tasks),
	}
}
