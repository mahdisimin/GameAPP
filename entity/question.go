package entity

type Question struct {
	ID         uint8
	CategoryID uint8
	Difficulty QuestionDifficulty
	Text       string
}

type QuestionDifficulty uint8

const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDifficultyMedium
	QuestionDifficultyHard
)
