package entity

type Player struct {
	ID         uint8
	UserID     uint8
	GameID     uint8
	AnswerList []int8
	Score      int8
	IsWinner   bool
}
