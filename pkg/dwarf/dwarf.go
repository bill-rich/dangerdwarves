package dwarf

import (
	"fmt"

	"github.com/bill-rich/dangerdwarves/pkg/task"
	"github.com/bill-rich/dangerdwarves/pkg/world"
)

type Dwarf struct {
	Name  string
	Age   int
	X     int
	Y     int
	Z     int
	Map   *world.LocalMap
	Tasks []task.Task
}

func (dwarf *Dwarf) ProcessTurn() {
	if len(dwarf.Tasks) > 0 {
		dwarf.ProcessNextTask()
		if len(dwarf.Tasks[0].Steps) == 0 {
			dwarf.Tasks = dwarf.Tasks[1:]
		}
	}
}

func (dwarf *Dwarf) ProcessNextTask() {
	step := dwarf.Tasks[0].Steps[0]
	switch step.StepType {
	case task.TaskWalk:
		dwarf.X = step.X
		dwarf.Y = step.Y
		dwarf.Z = step.Z
		fmt.Printf("%s moved to %d, %d, %d\n", dwarf.Name, dwarf.X, dwarf.Y, dwarf.Z)
	}
	dwarf.Tasks[0].Steps = dwarf.Tasks[0].Steps[1:]
}
