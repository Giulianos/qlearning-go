package qlearning

import (
	"testing"
)

func Test_PerformAction(t *testing.T) {
	e, a := GenerateEnvironment(
		Vec2D{1, 1},   // Agent position
		Vec2D{0, -1},  // Agent direction (up)
		Vec2D{9, 9},   // Goal position
		Vec2D{10, 10}, // Environment size
	)

	// Check that we can move upwards
	a.PerformAction(Forward)
	if e.agentPos.Y != 0 {
		t.Error("new agent position should be (1,0), got: ", e.agentPos)
	}

	// Check that now we can't move upwards
	a.PerformAction(Forward)
	if e.agentPos.Y != 0 {
		t.Error("agent position shouldn't change when hitting a boundary")
	}

	// Check that after rotating agent
	// moves to the right
	a.PerformAction(RotateR)
	a.PerformAction(Forward)
	if e.agentPos.Y != 0 || e.agentPos.X != 2 {
		t.Error("new agent position should be (2,0), got: ", e.agentPos)
	}

	// Check that it can move left
	a.PerformAction(RotateL)
	a.PerformAction(RotateL)
	a.PerformAction(Forward)
	if e.agentPos.Y != 0 || e.agentPos.X != 1 {
		t.Error("new agent position should be (1,0), got: ", e.agentPos)
	}

}
