package db

import (
	"github.com/Jordank321/pihouse/data"

	"github.com/jinzhu/gorm"
)

// NodeRepository is the data repository for Node readings
type NodeRepository interface {
	GetNodeByName(name string) *data.Node
	GetAllNodes() []*data.Node
	AddNode(node *data.Node)
}

// SQLNodeRepository is the MSSQL implimentation of NodeRepository
type SQLNodeRepository struct {
	Connection *gorm.DB
}

func (repository *SQLNodeRepository) Close() {
	repository.Connection.Close()
}

func (repository *SQLNodeRepository) GetNodeByName(name string) *data.Node {
	var node data.Node
	var count int
	repository.Connection.Where("name = ?", name).Find(&[]*data.Node{}).Count(&count)
	if count == 0 {
		return nil
	}
	repository.Connection.Where("name = ?", name).First(&node)
	return &node
}

func (repository *SQLNodeRepository) GetAllNodes() []*data.Node {
	var nodes []*data.Node
	if err := repository.Connection.Find(&nodes).Error; err != nil {
		panic(err)
	}
	return nodes
}

func (repository *SQLNodeRepository) AddNode(node *data.Node) {
	repository.Connection.Create(node)
}
