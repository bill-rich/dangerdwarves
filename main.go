package main

import (
	"time"

	"github.com/bill-rich/dangerdwarves/pkg/dwarf"
	"github.com/bill-rich/dangerdwarves/pkg/task"
	"github.com/bill-rich/dangerdwarves/pkg/world"
)

func main() {
	localMap := world.LocalMap{}
	localMap.GenerateLocalMap(10, 10, 10)
	steps := []task.Step{}
	steps = append(steps, task.Step{
		StepType: 0,
		X:        4,
		Y:        4,
		Z:        4,
	})
	steps = append(steps, task.Step{
		StepType: 0,
		X:        4,
		Y:        5,
		Z:        4,
	})
	steps = append(steps, task.Step{
		StepType: 0,
		X:        4,
		Y:        6,
		Z:        4,
	})

	moveTask := task.Task{
		TaskType: task.TaskWalk,
		Steps:    steps,
	}
	unit := dwarf.Dwarf{
		Name:  "Roger",
		Age:   100,
		X:     4,
		Y:     4,
		Z:     4,
		Map:   &localMap,
		Tasks: []task.Task{moveTask},
	}
	for run := true; run; run = true {
		unit.ProcessTurn()
		time.Sleep(time.Second * 1)
	}
}
