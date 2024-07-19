package counter

import (
	"Homework8/generator"
)

type TotalScore struct {
	ScorePerRound []int
	TotalScore    int
	AnswerHistory []int
}

func (t *TotalScore) Scores(round generator.Rounds, answer int) {
	t.AnswerHistory = append(t.AnswerHistory, answer)
	if answer == round.RightAnswer {
		t.ScorePerRound = append(t.ScorePerRound, 5)
	} else {
		t.ScorePerRound = append(t.ScorePerRound, 0)
	}
	t.TotalScore = 0
	for _, val := range t.ScorePerRound {
		t.TotalScore += val
	}
}
