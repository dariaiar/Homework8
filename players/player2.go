package players

import (
	"Homework8/generator"
	"fmt"
)

func GamePlayer2(c generator.Channels) {

	round1 := <-c.QuestionForPlayer

	//fmt.Printf("TEST print question %v", round1)
	for _, question := range round1.Question {
		fmt.Printf("\nPlayer 2. Your question is : %v\n", question)
		//break??
	}

	fmt.Printf("Player 2. Choose your answer:")

	for _, answer := range round1.AnswerVariants {
		fmt.Printf("\n %v   ", answer)
		//break??
	}

	var player2answer1 int
	fmt.Scan(&player2answer1)
	//time.Sleep(1 * time.Second)
	c.Answer <- player2answer1

}
