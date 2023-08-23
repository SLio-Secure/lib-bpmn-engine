package bpmn_engine

import "github.com/github.com/SLio-Secure/lib-bpmn-engine/pkg/spec/BPMN20"

func (state *BpmnEngineState) handleUserTask(process *ProcessInfo, instance *processInstanceInfo, element *BPMN20.TaskElement) bool {
	// TODO consider different handlers, since Service Tasks are different in their definition than user tasks
	return state.handleServiceTask(process, instance, element)
}
