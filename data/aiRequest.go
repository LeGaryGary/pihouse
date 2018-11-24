package data

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type AIRequest struct {
	gorm.Model
	Text    string
	Intents []Intent
	Node    *Node
	NodeID  uint `gorm:"index"`
}

type Intent struct {
	Value       string
	Confidence  decimal.Decimal `sql:"type:decimal(4,2)"`
	AIRequest   *AIRequest
	AIRequestID uint `gorm:"index"`
}
