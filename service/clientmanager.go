package service

import (
	"sync"
)

type ClientManager struct {
	ClientMap     map[string]*Client
	ClientMapLock sync.RWMutex

	ConnectChan    chan *Client
	DisconnectChan chan *Client

	Group     map[string]map[string]*Client
	GroupLock sync.RWMutex
}

var Manager = NewClientManager()

func NewClientManager() *ClientManager {
	return &ClientManager{
		ClientMap:      make(map[string]*Client),
		ConnectChan:    make(chan *Client, 1000),
		DisconnectChan: make(chan *Client, 1000),
		Group:          make(map[string]map[string]*Client),
	}
}

func (m *ClientManager) Start() {
	for {
		select {
		case client := <-m.ConnectChan:
			m.Connect(client)
		case client := <-m.DisconnectChan:
			m.Disconnect(client)
		}
	}
}

func (m *ClientManager) Connect(c *Client) {
	m.ClientMapLock.Lock()
	defer m.ClientMapLock.Unlock()

	m.ClientMap[c.ClientId] = c
	m.addClient2Group(c)
	go c.WritePump()
}

func (m *ClientManager) Disconnect(c *Client) {
	m.ClientMapLock.Lock()
	defer m.ClientMapLock.Unlock()

	m.deleteClientFromGroup(c)
	delete(m.ClientMap, c.ClientId)
	c.Conn.Close()
	c = nil
}

func (m *ClientManager) addClient2Group(c *Client) {
	m.GroupLock.Lock()
	defer m.GroupLock.Unlock()

	for _, v := range c.GroupList {
		if _, ok := m.Group[v]; ok == false {
			m.Group[v] = make(map[string]*Client)
		}
		m.Group[v][c.ClientId] = c
	}
}

func (m *ClientManager) deleteClientFromGroup(c *Client) {
	m.GroupLock.Lock()
	defer m.GroupLock.Unlock()

	for _, v := range c.GroupList {
		delete(m.Group[v], c.ClientId)
	}
}
