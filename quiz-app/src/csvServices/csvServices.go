package csvServices

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"quiz-app/src/models"
	"strconv"
)

func UpdateScoresCsv() error {

	playerRows := [][]string{}
	for _, player := range models.Players {
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

func ReadScoresCsv() ([]models.PlayerInfo, error) {
	csvFile, err := os.OpenFile("scores.csv", os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Error opening CSV file: %w", err)
	}
	csvReader := csv.NewReader(csvFile)
	players := []models.PlayerInfo{}

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

		playerInfo := models.PlayerInfo{
			ID:    playerID,
			Score: playerScoreToString,
		}

		players = append(players, playerInfo)
	}

	return players, nil
}

func CalculateTopScorePercentage(userScore float64, players []models.PlayerInfo) float64 {
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
