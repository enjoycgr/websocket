package service

import (
	"sync"
)

type GroupManager struct {
	Group map[string]map[string]*Client
	Lock  sync.RWMutex
}

var G = NewGroupManager()

func NewGroupManager() *GroupManager {
	return &GroupManager{}
}
