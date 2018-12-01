package db

import (
	"log"

	"github.com/Jordank321/pihouse/data"
	"github.com/jinzhu/gorm"
)

type AIRepository interface {
	NewWitAIOutcome(request *data.AIRequest)
	FindActions(intentValue string) []data.ActionMaping
}

type SQLAIRespository struct {
	Connection *gorm.DB
}

func (repository *SQLAIRespository) Close() {
	repository.Connection.Close()
}

func (repository *SQLAIRespository) NewWitAIOutcome(request *data.AIRequest) {
	repository.Connection.Create(request)
}

func (repository *SQLAIRespository) FindActions(intentValue string) []data.ActionMaping {
	actions := []data.ActionMaping{}
	sql := repository.Connection.Where("intent_value LIKE '%" + intentValue + "%'")
	sql.Find(&actions)
	log.Printf("Found %d actions", len(actions))
	return actions
}
