package generator

import "time"

type Rounds struct {
	Question       map[int]string
	AnswerVariants map[int]string
	RightAnswer    int
}

var round1 Rounds = Rounds{
	Question:       map[int]string{1: "How many days are in the week?"},
	AnswerVariants: map[int]string{1: "1) 7 days", 2: "2) 10 days", 3: "3) 100 days", 4: "4) 17 days"},
	RightAnswer:    1,
}

var round2 Rounds = Rounds{
	Question:       map[int]string{2: "What is the capital of Greate Britain?"},
	AnswerVariants: map[int]string{1: "1) London", 2: "2) Paris", 3: "3) Kyiv", 4: "4) Athens"},
	RightAnswer:    1,
}

var round3 Rounds = Rounds{
	Question:       map[int]string{1: "To be or not to be?"},
	AnswerVariants: map[int]string{1: "1) to be, but after weekend", 2: "2) to be", 3: "3) not to be", 4: "4) b2b "},
	RightAnswer:    2,
}

type Channels struct {
	QuestionForPlayer chan Rounds
	QuestionForScores chan Rounds
	Answer            chan int
}

func NewRound(c Channels, rou int) {

	whichround := rou
	switch whichround {
	case 1:
		c.QuestionForPlayer <- round1
		c.QuestionForScores <- round1
		time.Sleep(10 * time.Second)
		close(c.QuestionForPlayer)
		close(c.QuestionForScores)
	case 2:
		c.QuestionForPlayer <- round2
		c.QuestionForScores <- round2
		time.Sleep(10 * time.Second)
		close(c.QuestionForPlayer)
		close(c.QuestionForScores)
	case 3:
		c.QuestionForPlayer <- round3
		c.QuestionForScores <- round3
		time.Sleep(10 * time.Second)
		close(c.QuestionForPlayer)
		close(c.QuestionForScores)
	}

}
