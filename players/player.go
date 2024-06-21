package players

import (
	"Homework8/generator"
	"context"
	"fmt"
	"math/rand"
	"time"
)

func GamePlayer(c generator.Channels, ctx context.Context, playerID int) {
	round := <-c.QuestionForPlayer
	//fmt.Printf("TEST print question %v", round1)
	for _, question := range round.Question {
		fmt.Printf("\nPlayer %v. Your question is : %v\n", playerID, question)
	}
	fmt.Printf("Player %v. Choose your answer:", playerID)
	for _, answer := range round.AnswerVariants {
		fmt.Printf("\n %v  ", answer)
	}

	select {
	case <-ctx.Done():
		fmt.Printf("\n Game over for player %v", playerID)
	default:
		fmt.Println("\n continue working...")
		playerAnswer := rand.Intn(4)
		c.Answer <- playerAnswer
		fmt.Printf("\nAnswer: %v", playerAnswer)
		time.Sleep(1 * time.Second)
	}
}
