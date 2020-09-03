package task

import "github.com/bill-rich/dangerdwarves/pkg/common"

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
	StepType   int
	Coordinate common.Coordinate
}
