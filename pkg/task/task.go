package task

const (
	TaskWalk = iota
)

type Queue struct {
	Queue []Task
}

type Task struct {
	TaskType int
	Steps    []Step
}

type Step struct {
	StepType int
	X        int
	Y        int
	Z        int
}
