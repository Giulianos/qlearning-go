package main

import (
	"fmt"
	"github.com/Giulianos/qlearning-go"
)

func main() {
	knowledge := qlearning.InitKnowledge()

	knowledge.Learn(
		func(s1, s2 qlearning.AgentState) float64 {
			if s2.Dist.Norm() < s1.Dist.Norm() {
				return 1.0
			} else {
				return -1.0
			}
		},
		1000,
		func(episode int) float64 { return 0.6 },
		1.0,
		func(episode int) float64 { return 0.1 },
	)

	// Example 1
	env, agent := qlearning.GenerateEnvironment(
		qlearning.Vec2D{5, 5},   // Agent position
		qlearning.Vec2D{0, -1},  // Agent direction (up)
		qlearning.Vec2D{9, 9},   // Goal position
		qlearning.Vec2D{10, 10}, // Environment size
	)

	// Print env
	fmt.Println(env)

	// Run agent (following knowledge policy)
	knowledge.Run(env, agent)
}
