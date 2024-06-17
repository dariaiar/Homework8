package players

import (
	"Homework8/generator"
	"fmt"
)

func GamePlayer3(c generator.Channels) {

	round1 := <-c.QuestionForPlayer

	//fmt.Printf("TEST print question %v", round1)
	for _, question := range round1.Question {
		fmt.Printf("\nPlayer 3. Your question is : %v\n", question)
		//break??
	}

	fmt.Printf("Player 3. Choose your answer:")

	for _, answer := range round1.AnswerVariants {
		fmt.Printf("\n %v   ", answer)
		//break??
	}

	var player1answer1 int
	fmt.Scan(&player1answer1)
	//time.Sleep(1 * time.Second)
	c.Answer <- player1answer1

}
