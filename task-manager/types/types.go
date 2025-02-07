package types

type TaskPriority uint
type TaskStatus uint
type Tasks map[uint]*Task
type Teams string;

type Task struct {
	Id uint
	Title    string
	TeamName Teams
	Priority TaskPriority
	Status   TaskStatus
}

type TaskRepository struct {
	AllTasks Tasks
}
