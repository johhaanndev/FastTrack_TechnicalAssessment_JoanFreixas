package models

type Question struct {
	Text    string   `json:"text"`
	Answers []string `json:"possibleAnswers"`
	Correct string   `json:"correctAnswer"`
}

type GetQuestionsReponse struct {
	Text    string   `json:"text"`
	Answers []string `json:"possibleAnswers"`
}

type PlayerInfo struct {
	ID    string `json:"id"`
	Score string `json:"score"`
}

type PlayerInfoCsv struct {
	Id    string `csv:"id"`
	Score string `csv:"score"`
}

func (q *Question) ToGetQuestionsResponse() GetQuestionsReponse {
	return GetQuestionsReponse{
		Text:    q.Text,
		Answers: q.Answers,
	}
}

var QuizQuestions = []Question{
	{
		Text:    "What is the largest planet in our solar system?",
		Answers: []string{"a. Earth", "b. Jupiter", "c. Mars"},
		Correct: "b",
	},
	{
		Text:    "What is the capital of France?",
		Answers: []string{"a. London", "b. Berlin", "c. Paris"},
		Correct: "c",
	},
	{
		Text:    "What is the tallest mountain in the world?",
		Answers: []string{"a. Mount Everest", "b. K2", "c. Mount Kilimanjaro"},
		Correct: "a",
	},
	{
		Text:    "How many colors are there in the rainbow?",
		Answers: []string{"a. 5", "b. 7", "c. 10"},
		Correct: "b",
	},
	{
		Text:    "What is the name of the world wide web inventor?",
		Answers: []string{"a. Bill Gates", "b. Steve Jobs", "c. Tim Berners-Lee"},
		Correct: "c",
	},
}

var Players = []PlayerInfo{}
