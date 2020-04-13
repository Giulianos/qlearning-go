package qlearning

type Agent struct {
	r   *Vec2D
	dir *Vec2D
	env *Environment
}

type AgentState struct {
	Dist Vec2D
	Dir  Vec2D
}

func NewAgent(r, dir Vec2D, env *Environment) Agent {
	return Agent{&r, &dir, env}
}

func (a *Agent) PerformAction(action Action) {
	switch action {
	case Forward:
		if a.env.CheckBoundary(*a.dir) {
			r := a.r.Subtract(*a.dir)
			a.r = &r
			a.env.NotifyAgentMove(*a.dir)
		}
		break
	case RotateR:
		dir := a.dir.RotateCW()
		a.dir = &dir
		break
	case RotateL:
		dir := a.dir.RotateCCW()
		a.dir = &dir
		break
	}
}

func (a Agent) GetState() AgentState {
	return AgentState{
		Dist: *a.r,
		Dir:  *a.dir,
	}
}
