package users

import (
	"Homework8/generator"
	"context"
	"fmt"
	"math/rand"
	"sync"
)

type Player struct {
	ID      int
	Channel chan generator.Rounds
}

type PlayerAnswer struct {
	PlayerID int
	Answer   int
}

func (p *Player) Play(ctx context.Context, wg *sync.WaitGroup, answerChan chan<- PlayerAnswer) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nGame over for player %v", p.ID)
			return
		case round, ok := <-p.Channel:
			if !ok {
				return
			}
			fmt.Printf("\nPlayer %v. Your question is : %v\n", p.ID, round.Question)
			for _, answer := range round.AnswerVariants {
				fmt.Printf("\n %v  ", answer)
			}
			fmt.Printf("\nPlayer %v. Choose your answer:", p.ID)
			playerAnswer := rand.Intn(4) + 1
			fmt.Printf("\nAnswer: %v", playerAnswer)
			answerChan <- PlayerAnswer{PlayerID: p.ID, Answer: playerAnswer}
		}
	}
}
