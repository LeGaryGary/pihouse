package db

import (
	"github.com/Jordank321/pihouse/data"
	"github.com/jinzhu/gorm"
)

type AIRepository interface {
	NewWitAIOutcome(request *data.AIRequest)
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
