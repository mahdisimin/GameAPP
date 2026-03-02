package entity

type Answer struct {
	ID         uint8
	Text       string
	QuestionID uint8
	Difficulty uint8
	IsCorrect  bool
}
