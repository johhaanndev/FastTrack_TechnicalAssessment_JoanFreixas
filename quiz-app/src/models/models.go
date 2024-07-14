package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

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

func UpdateScoresCsv() error {

	playerRows := [][]string{}
	for _, player := range Players {
		playerRow := []string{player.ID, player.Score}
		playerRows = append(playerRows, playerRow)
	}

	csvFile, err := os.OpenFile("scores.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(csvFile)

	csvWriter.Write(playerRows[len(playerRows)-1])

	csvWriter.Flush()
	csvFile.Close()

	return nil
}

func ReadScoresCsv() ([]PlayerInfo, error) {
	csvFile, err := os.OpenFile("scores.csv", os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Error opening CSV file: %w", err)
	}
	csvReader := csv.NewReader(csvFile)
	players := []PlayerInfo{}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Error reading CSV file: %w", err)
		}

		playerID := record[0]
		playerScore, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, fmt.Errorf("Error parsing player score: %w", err)
		}
		playerScoreToString := fmt.Sprintf("%.2f", playerScore)

		playerInfo := PlayerInfo{
			ID:    playerID,
			Score: playerScoreToString,
		}

		players = append(players, playerInfo)
	}

	return players, nil
}

func CalculateTopScorePercentage(userScore float64, players []PlayerInfo) float64 {
	position := 0

	for _, player := range players {
		playerScore, err := strconv.ParseFloat(player.Score, 64)
		if err != nil {
			continue
		}

		if playerScore <= userScore {
			position++
		}
	}

	totalPlayers := len(players)
	if totalPlayers == 0 {
		return 0.0
	}
	percentage := float64(position) / float64(totalPlayers) * 100.0

	return percentage
}

func (q *Question) ToResponseQuestion() GetQuestionsReponse {
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
