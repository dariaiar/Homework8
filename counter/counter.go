package counter

import (
	"Homework8/generator"
	"fmt"
)

type TotalScore struct {
	ID            int
	ScorePerRound []int
	TotalScore    int
	Answer        int
	AnswerHistory []int
}

var Player1 = TotalScore{
	ID:            82,
	ScorePerRound: []int{},
	TotalScore:    0,
	Answer:        0,
	AnswerHistory: []int{},
}

var Player2 = TotalScore{
	ID:            107,
	ScorePerRound: []int{},
	TotalScore:    0,
	Answer:        0,
	AnswerHistory: []int{},
}

var Player3 = TotalScore{
	ID:            315,
	ScorePerRound: []int{},
	TotalScore:    0,
	Answer:        0,
	AnswerHistory: []int{},
}

func (t *TotalScore) Scores(round generator.Rounds, answer int) TotalScore {

	t.AnswerHistory = append(t.AnswerHistory, answer)

	var sum int
	if answer == round.RightAnswer {
		t.ScorePerRound = append(t.ScorePerRound, 5)
		for _, val := range t.ScorePerRound {
			//fmt.Println(val)
			sum += val
		}
		fmt.Printf("Correct answer. You gained 5 points. Your total score is: %v", sum)
	} else {
		t.ScorePerRound = append(t.ScorePerRound, 0)
		for _, val := range t.ScorePerRound {
			//fmt.Println(val)
			sum += val
		}
		fmt.Printf("Wrong answer. You gained 0 points. Your total score is: %v", sum)

	}
	t.TotalScore = sum
	return *t
}
