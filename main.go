package main

import (
	"Homework8/counter"
	"Homework8/generator"
	"Homework8/users"
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	fmt.Println("---------------START GAME---------------")
	numPlayers := 3
	playersScores := make([]*counter.TotalScore, numPlayers)
	for i := 0; i < numPlayers; i++ {
		playersScores[i] = &counter.TotalScore{}
	}
	players := make([]*users.Player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		players[i] = &users.Player{
			ID:      i,
			Channel: make(chan generator.Rounds),
		}
	}
	var wg sync.WaitGroup
	answerChan := make(chan users.PlayerAnswer)
	for i := range players {
		wg.Add(1)
		go players[i].Play(ctx, &wg, answerChan)
	}
	go func() {
		wg.Wait()
		close(answerChan)
	}()
	rounds := []generator.Rounds{generator.Round1, generator.Round2, generator.Round3}
	for round := 1; round <= 3; round++ {
		fmt.Printf("\n----------ROUND %v----------\n", round)
		for i := range players {
			players[i].Channel <- rounds[round-1]
		}
		for i := 0; i < numPlayers; i++ {
			playerAnswer := <-answerChan
			playersScores[playerAnswer.PlayerID].Scores(rounds[round-1], playerAnswer.Answer)
			fmt.Printf("\nscorePerRound: %v;\ntotalScore: %v; \nanswerHistory: %v;\n",
				playersScores[playerAnswer.PlayerID].ScorePerRound,
				playersScores[playerAnswer.PlayerID].TotalScore,
				playersScores[playerAnswer.PlayerID].AnswerHistory)
		}
		time.Sleep(2 * time.Second)
	}
	for i := range players {
		close(players[i].Channel)
	}
	//---- WINNER ----
	fmt.Printf("\nTotal scores of players:\n")
	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Player %v: %v\n", i, playersScores[i].TotalScore)
	}
}
