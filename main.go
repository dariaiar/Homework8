package main

import (
	"Homework8/counter"
	"Homework8/generator"
	"Homework8/players"
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	fmt.Println("---------------START GAME---------------")
	numPlayers := 3
	playersScores := make([]counter.TotalScore, numPlayers)
	for round := 1; round <= 3; round++ {
		fmt.Printf("\n----------ROUND %v----------\n", round)

		for i := 0; i < numPlayers; i++ {
			chanRound := generator.Channels{
				QuestionForPlayer: make(chan generator.Rounds),
				QuestionForScores: make(chan generator.Rounds),
				Answer:            make(chan int),
			}
			go generator.NewRound(chanRound, round)
			go players.GamePlayer(chanRound, ctx, i)
			roundData := <-chanRound.QuestionForScores
			answer := <-chanRound.Answer
			playersScores[i] = playersScores[i].Scores(roundData, answer)
			fmt.Printf("\nscorePerRound: %v;\ntotalScore: %v; \nanswerHistory: %v;\n", playersScores[i].ScorePerRound, playersScores[i].TotalScore, playersScores[i].AnswerHistory)
			time.Sleep(2 * time.Second)
		}
	}
	//---- WINNER ----
	fmt.Printf("\nTotal scores of players %v", playersScores)
}
