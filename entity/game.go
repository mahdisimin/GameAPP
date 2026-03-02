package entity

type Game struct {
	ID           uint8
	CategoryID   uint8
	QuestionList []int8
	difficulty   uint8
}
