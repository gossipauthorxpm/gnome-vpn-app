package connection

import (
	"container/list"
)

type ConnectionList struct {
	list *list.List
}

type ConnectionListUnit struct {
	name    string
	country string
}

func CreateConnectionList() *ConnectionList {
	return &ConnectionList{}
}
