package dwarf

import (
	"fmt"

	"github.com/bill-rich/dangerdwarves/pkg/common"
	"github.com/bill-rich/dangerdwarves/pkg/task"
	"github.com/bill-rich/dangerdwarves/pkg/world"
)

type Dwarf struct {
	Name       string
	Age        int
	Coordinate common.Coordinate
	Map        *world.LocalMap
	Tasks      []task.Task
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
		dwarf.Coordinate = step.Coordinate
		fmt.Printf("%s moved to %#v\n", dwarf.Name, dwarf.Coordinate)
	}
	dwarf.Tasks[0].Steps = dwarf.Tasks[0].Steps[1:]
}
