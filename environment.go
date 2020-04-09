package qlearning

type Environment struct {
	// This are used for boundary
	// checking and rendering
	goalPos  *Vec2D
	agentPos *Vec2D

	// Dimensions
	dim *Vec2D
}

func GenerateEnvironment(agentPos, agentDir, goalPos, dim Vec2D) (*Environment, *Agent) {
	e := Environment{
		goalPos:  &goalPos,
		agentPos: &agentPos,

		dim: &dim,
	}

	a := NewAgent(goalPos.Subtract(agentPos), agentDir, &e)

	return &e, &a
}

func (e *Environment) CheckBoundary(movementDir Vec2D) bool {
	newPos := e.agentPos.Add(movementDir)
	return newPos.X < e.dim.X && newPos.X >= 0 && newPos.Y < e.dim.Y && newPos.Y >= 0
}

func (e *Environment) NotifyAgentMove(movementDir Vec2D) {
	agentPos := e.agentPos.Add(movementDir)
	e.agentPos = &agentPos
}
