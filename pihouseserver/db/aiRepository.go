package db

import (
	"github.com/Jordank321/pihouse/data"
	"github.com/jinzhu/gorm"
)

type AIRepository interface {
	NewWitAIOutcome(request *data.AIRequest)
	FindAction(intentValue string) *data.Action
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

func (repository *SQLAIRespository) FindAction(intentValue string) *data.Action {
	count := 0
	repository.Connection.Where("intent_value = " + intentValue).Find(&[]*data.ActionMaping{}).Count(&count)
	if count > 0 {
		action := data.ActionMaping{}
		repository.Connection.Where("intent_value = " + intentValue).First(&action)
		return &action.Action
	}

	return nil
}
