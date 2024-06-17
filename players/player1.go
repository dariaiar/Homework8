package players

import (
	"Homework8/generator"
	"context"
	"fmt"
	"time"
)

func GamePlayer1(c generator.Channels, ctx context.Context) {
	round1 := <-c.QuestionForPlayer
	//fmt.Printf("TEST print question %v", round1)
	for _, question := range round1.Question {
		fmt.Printf("\nPlayer 1. Your question is : %v\n", question)
	}
	fmt.Printf("Player 1. Choose your answer:")
	for _, answer := range round1.AnswerVariants {
		fmt.Printf("\n %v   ", answer)
	}

	//???????????????
	newCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var player1answer1 int
	fmt.Scan(&player1answer1)
	//time.Sleep(1 * time.Second)

	for {
		select {
		case <-newCtx.Done():
			fmt.Printf("ctx done:")

		default:
			fmt.Println("continue working...")
			c.Answer <- player1answer1
			time.Sleep(1 * time.Second)
		}
	}

}
