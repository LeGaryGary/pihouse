package data

import (
	"github.com/jinzhu/gorm"
)

type ActionMaping struct {
	gorm.Model
	Action      Action
	IntentValue string
}
