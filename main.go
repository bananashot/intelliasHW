package main

import (
	"strings"
)

func getWordScore(scoreMap map[rune]int, word string) int {
	var totalScore int
	for _, v := range word {
		totalScore += scoreMap[v]
	}
	return totalScore
}

func High(s string) string {

	alphabetScore := make(map[rune]int)
	score := 1

	for i := 'a'; i <= 'z'; i++ {
		alphabetScore[i] = score
		score++
	}

	wordAray := strings.Split(s, " ")
	var highestScore int
	var earliestHighScoreIndex int

	for i, v := range wordAray {

		if highestScore < getWordScore(alphabetScore, v) {
			highestScore = getWordScore(alphabetScore, v)
			earliestHighScoreIndex = i
		}
	}
	return wordAray[earliestHighScoreIndex]
}

func main() {
	High("man i need a taxi up to ubud")
}
