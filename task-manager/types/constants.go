package types

const (
	Low TaskPriority = iota
	Medium
	High
	Unassigned TaskStatus = iota
	Assigned
	Inprogress
	Inreview
	Completed
	Frontend Teams = "Frontend"
	Backend Teams = "Backend"
)