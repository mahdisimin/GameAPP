package entity

type Game struct {
	ID           uint8
	CategoryID   uint8
	QuestionList []int8
	Difficulty   GameDifficulty
}

type GameDifficulty uint8

const (
	GameDifficultyEasy GameDifficulty = iota + 1
	GameDifficultyMedium
	GameDifficultyHard
)
