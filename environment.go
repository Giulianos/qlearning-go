package qlearning

type Environment struct {
	// This are used for boundary
	// checking and rendering
	goalPos  *Vec2D
	agentPos *Vec2D
	agent *Agent

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
	e.agent = &a

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

func (e Environment) GoalAchived() bool {
	return e.agentPos.X == e.goalPos.X && e.agentPos.Y == e.goalPos.Y
}

func (e Environment) String() string {
	s := ""
	agentDir := e.agent.dir

	for y := 0; y < e.dim.Y; y++ {
		for x := 0; x < e.dim.X; x++ {
			if x == e.goalPos.X && y == e.goalPos.Y {
				s += "X "
			} else if x == e.agentPos.X && y == e.agentPos.Y {
				switch *agentDir {
				case Vec2D{0, -1}: s += "▲ "
				case Vec2D{0, 1}: s += "▼ "
				case Vec2D{-1, 0}: s += "◄ "
				case Vec2D{1, 0}: s += "► "
				}
			} else {
				s += "- "
			}
		}
		s += "\n"
	}

	return s
}
