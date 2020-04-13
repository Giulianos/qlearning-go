package qlearning

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

type EpisodicParam func(episode int) float64
type RewardFunc func(s1, s2 AgentState) float64

type QLearning struct {
	Q map[AgentState]map[Action]float64
}

func InitKnowledge() QLearning {
	return QLearning{
		Q: make(map[AgentState]map[Action]float64, 0),
	}
}

func (ql *QLearning) Learn(reward RewardFunc, maxEpisodes int, alpha EpisodicParam, discountRate float64, epsilon EpisodicParam) {

	for episode := 0; episode < maxEpisodes; episode++ {
		// TODO: this should be randomized on each episode
		env, agent := GenerateEnvironment(
			Vec2D{1, 1},   // Agent position
			Vec2D{0, -1},  // Agent direction (up)
			Vec2D{9, 9},   // Goal position
			Vec2D{10, 10}, // Environment size
		)

		for !env.GoalAchived() {

			s := agent.GetState()
			// Choose action from Q using epsilon-greedy
			a := ql.getPolicy(s, epsilon(episode))

			// Perform action
			agent.PerformAction(a)

			// Observe state
			sPrime := agent.GetState()

			// Get reward
			r := reward(s, sPrime)

			// Update state-action
			bestNextAction := ql.getPolicy(sPrime, 0.0)
			tdTarget := r + discountRate*ql.Q[sPrime][bestNextAction]
			tdDelta := tdTarget - ql.Q[s][a]
			ql.Q[s][a] += alpha(episode) * tdDelta
		}

		log.Println("Finished episode: ", episode)
	}
}

func bestAction(actions map[Action]float64) Action {
	var bestAction Action
	bestValue := -1 * math.MaxFloat64
	for a, v := range actions {
		if v > bestValue {
			bestAction = a
			bestValue = v
		}
	}

	return bestAction
}

func (ql QLearning) getPolicy(state AgentState, epsilon float64) Action {
	// Make sure state exists in Q
	_, exists := ql.Q[state]
	if !exists {
		ql.Q[state] = make(map[Action]float64, 4)
		ql.Q[state][Forward] = 0
		ql.Q[state][RotateL] = 0
		ql.Q[state][RotateR] = 0
	}

	// Return best action with probability (1-epsilon)
	if rand.Float64() >= epsilon {
		return bestAction(ql.Q[state])
	}

	// Return other random action
	return Action(rand.Intn(len(ql.Q[state])))
}

func (ql QLearning) Run(env *Environment, agent *Agent) {
	for !env.GoalAchived() {
		s := agent.GetState()

		// Choose action from Q (greedy)
		a := ql.getPolicy(s, 0)

		agent.PerformAction(a)
		fmt.Println(env)
	}
}