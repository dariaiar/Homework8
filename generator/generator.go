package generator

type Rounds struct {
	Question       map[int]string
	AnswerVariants map[int]string
	RightAnswer    int
}

var Round1 Rounds = Rounds{
	Question:       map[int]string{1: "How many days are in the week?"},
	AnswerVariants: map[int]string{1: "1) 7 days", 2: "2) 10 days", 3: "3) 100 days", 4: "4) 17 days"},
	RightAnswer:    1,
}

var Round2 Rounds = Rounds{
	Question:       map[int]string{2: "What is the capital of Greate Britain?"},
	AnswerVariants: map[int]string{1: "1) London", 2: "2) Paris", 3: "3) Kyiv", 4: "4) Athens"},
	RightAnswer:    1,
}

var Round3 Rounds = Rounds{
	Question:       map[int]string{1: "To be or not to be?"},
	AnswerVariants: map[int]string{1: "1) to be, but after weekend", 2: "2) to be", 3: "3) not to be", 4: "4) b2b "},
	RightAnswer:    2,
}
