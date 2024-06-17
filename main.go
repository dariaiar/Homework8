package main

import (
	"Homework8/counter"
	"Homework8/generator"
	"Homework8/players"
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("---------------START GAME---------------")
	var player1Score counter.TotalScore
	var player2Score counter.TotalScore
	var player3Score counter.TotalScore
	//ROUND1
	fmt.Println("\n----------ROUND 1----------")
	const round1 = 1
	//-----Player1-----
	var chanPlayer1Round1 = generator.Channels{
		QuestionForPlayer: make(chan generator.Rounds),
		QuestionForScores: make(chan generator.Rounds),
		Answer:            make(chan int),
	}
	go generator.NewRound(chanPlayer1Round1, round1)
	go players.GamePlayer1(chanPlayer1Round1, ctx)
	round11 := <-chanPlayer1Round1.QuestionForScores
	answer11 := <-chanPlayer1Round1.Answer
	//var player1Score counter.TotalScore
	player1Score = counter.Player1.Scores(round11, answer11)
	fmt.Printf("\nID: %v;\nscorePerRound: %v;\ntotalScore: %v; \nanswer:  %v;\nanswerHistory: %v;\n", player1Score.ID, player1Score.ScorePerRound, player1Score.TotalScore, player1Score.Answer, player1Score.AnswerHistory)
	time.Sleep(2 * time.Second)

	//-----Player2-----3
	var chanPlayer2Round1 = generator.Channels{
		QuestionForPlayer: make(chan generator.Rounds),
		QuestionForScores: make(chan generator.Rounds),
		Answer:            make(chan int),
	}
	go generator.NewRound(chanPlayer2Round1, round1)
	go players.GamePlayer2(chanPlayer2Round1)
	round12 := <-chanPlayer2Round1.QuestionForScores
	answer12 := <-chanPlayer2Round1.Answer
	//var player2Score counter.TotalScore
	player2Score = counter.Player2.Scores(round12, answer12)
	fmt.Printf("\nID: %v;\nscorePerRound: %v;\ntotalScore: %v; \nanswer:  %v;\nanswerHistory: %v;\n", player2Score.ID, player2Score.ScorePerRound, player2Score.TotalScore, player2Score.Answer, player1Score.AnswerHistory)
	time.Sleep(2 * time.Second)

	//-----Player3-----
	var chanPlayer3Round1 = generator.Channels{
		QuestionForPlayer: make(chan generator.Rounds),
		QuestionForScores: make(chan generator.Rounds),
		Answer:            make(chan int),
	}
	go generator.NewRound(chanPlayer3Round1, round1)
	go players.GamePlayer3(chanPlayer3Round1)
	round13 := <-chanPlayer3Round1.QuestionForScores
	answer13 := <-chanPlayer3Round1.Answer
	//var player3Score counter.TotalScore
	player3Score = counter.Player3.Scores(round13, answer13)
	fmt.Printf("\nID: %v;\nscorePerRound: %v;\ntotalScore: %v; \nanswer:  %v;\nanswerHistory: %v;\n", player3Score.ID, player3Score.ScorePerRound, player3Score.TotalScore, player3Score.Answer, player3Score.AnswerHistory)
	time.Sleep(2 * time.Second)

	//ROUND2
	fmt.Println("\n----------ROUND 2----------")
	const round2 = 2
	//-----Player1-----
	var chanPlayer1Round2 = generator.Channels{
		QuestionForPlayer: make(chan generator.Rounds),
		QuestionForScores: make(chan generator.Rounds),
		Answer:            make(chan int),
	}
	go generator.NewRound(chanPlayer1Round2, round2)
	go players.GamePlayer1(chanPlayer1Round2, ctx)
	round21 := <-chanPlayer1Round2.QuestionForScores
	answer21 := <-chanPlayer1Round2.Answer
	//var player1Score counter.TotalScore
	player1Score = counter.Player1.Scores(round21, answer21)
	fmt.Printf("\nID: %v;\nscorePerRound: %v;\ntotalScore: %v; \nanswer:  %v;\nanswerHistory: %v;\n", player1Score.ID, player1Score.ScorePerRound, player1Score.TotalScore, player1Score.Answer, player1Score.AnswerHistory)
	time.Sleep(2 * time.Second)

	//-----Player2-----
	var chanPlayer2Round2 = generator.Channels{
		QuestionForPlayer: make(chan generator.Rounds),
		QuestionForScores: make(chan generator.Rounds),
		Answer:            make(chan int),
	}
	go generator.NewRound(chanPlayer2Round2, round2)
	go players.GamePlayer2(chanPlayer2Round2)
	round22 := <-chanPlayer2Round2.QuestionForScores
	answer22 := <-chanPlayer2Round2.Answer
	//var player1Score counter.TotalScore
	player2Score = counter.Player2.Scores(round22, answer22)
	fmt.Printf("\nID: %v;\nscorePerRound: %v;\ntotalScore: %v; \nanswer:  %v;\nanswerHistory: %v;\n", player2Score.ID, player2Score.ScorePerRound, player2Score.TotalScore, player2Score.Answer, player2Score.AnswerHistory)
	time.Sleep(2 * time.Second)

	//-----Player3-----
	var chanPlayer3Round2 = generator.Channels{
		QuestionForPlayer: make(chan generator.Rounds),
		QuestionForScores: make(chan generator.Rounds),
		Answer:            make(chan int),
	}
	go generator.NewRound(chanPlayer3Round2, round2)
	go players.GamePlayer3(chanPlayer3Round2)
	round23 := <-chanPlayer3Round2.QuestionForScores
	answer23 := <-chanPlayer3Round2.Answer
	//var player1Score counter.TotalScore
	player3Score = counter.Player3.Scores(round23, answer23)
	fmt.Printf("\nID: %v;\nscorePerRound: %v;\ntotalScore: %v; \nanswer:  %v;\nanswerHistory: %v;\n", player3Score.ID, player3Score.ScorePerRound, player3Score.TotalScore, player3Score.Answer, player3Score.AnswerHistory)
	time.Sleep(2 * time.Second)

	//ROUND2
	fmt.Println("\n----------ROUND 3----------")
	const round3 = 3
	//-----Player1-----
	var chanPlayer1Round3 = generator.Channels{
		QuestionForPlayer: make(chan generator.Rounds),
		QuestionForScores: make(chan generator.Rounds),
		Answer:            make(chan int),
	}
	go generator.NewRound(chanPlayer1Round3, round3)
	go players.GamePlayer1(chanPlayer1Round3, ctx)
	round31 := <-chanPlayer1Round3.QuestionForScores
	answer31 := <-chanPlayer1Round3.Answer
	//var player1Score counter.TotalScore
	player1Score = counter.Player1.Scores(round31, answer31)
	fmt.Printf("\nID: %v;\nscorePerRound: %v;\ntotalScore: %v; \nanswer:  %v;\nanswerHistory: %v;\n", player1Score.ID, player1Score.ScorePerRound, player1Score.TotalScore, player1Score.Answer, player1Score.AnswerHistory)
	time.Sleep(2 * time.Second)

	//-----Player2-----
	var chanPlayer2Round3 = generator.Channels{
		QuestionForPlayer: make(chan generator.Rounds),
		QuestionForScores: make(chan generator.Rounds),
		Answer:            make(chan int),
	}
	go generator.NewRound(chanPlayer2Round3, round3)
	go players.GamePlayer2(chanPlayer2Round3)
	round32 := <-chanPlayer2Round3.QuestionForScores
	answer32 := <-chanPlayer2Round3.Answer
	//var player1Score counter.TotalScore
	player2Score = counter.Player3.Scores(round32, answer32)
	fmt.Printf("\nID: %v;\nscorePerRound: %v;\ntotalScore: %v; \nanswer:  %v;\nanswerHistory: %v;\n", player2Score.ID, player2Score.ScorePerRound, player2Score.TotalScore, player2Score.Answer, player2Score.AnswerHistory)
	time.Sleep(2 * time.Second)

	//-----Player3-----
	var chanPlayer3Round3 = generator.Channels{
		QuestionForPlayer: make(chan generator.Rounds),
		QuestionForScores: make(chan generator.Rounds),
		Answer:            make(chan int),
	}
	go generator.NewRound(chanPlayer3Round3, round3)
	go players.GamePlayer3(chanPlayer3Round3)
	round33 := <-chanPlayer3Round3.QuestionForScores
	answer33 := <-chanPlayer3Round3.Answer
	//var player1Score counter.TotalScore
	player3Score = counter.Player3.Scores(round33, answer33)
	fmt.Printf("\nID: %v;\nscorePerRound: %v;\ntotalScore: %v; \nanswer:  %v;\nanswerHistory: %v;\n", player3Score.ID, player3Score.ScorePerRound, player3Score.TotalScore, player3Score.Answer, player3Score.AnswerHistory)
	time.Sleep(2 * time.Second)

	//---- WINNER ----
	fmt.Printf("Total score of Player1 is %v;\nTotal score of Player2 is %v;\nTotal score of Player3 is %v.", player1Score.TotalScore, player2Score.TotalScore, player3Score.TotalScore)

}
