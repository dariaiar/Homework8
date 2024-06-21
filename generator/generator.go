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

func NewChannels() Channels {
	return Channels{
		QuestionForPlayer: make(chan Rounds),
		QuestionForScores: make(chan Rounds),
		Answer:            make(chan int),
	}
}

func NewRound(c Channels, roundNum int) {
	var round Rounds

	switch roundNum {
	case 1:
		round = round1
	case 2:
		round = round2
	case 3:
		round = round3
	}

	c.QuestionForPlayer <- round
	c.QuestionForScores <- round
	time.Sleep(10 * time.Second)
	close(c.QuestionForPlayer)
	close(c.QuestionForScores)
}
