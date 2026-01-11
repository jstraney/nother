// Node registry for the editor
package registry

import "github.com/jstraney/nother/pkg/node"

type Store struct {
	NodeTypes map[string]Entry
}

type Factory func() node.Node

func NewStore() *Store {
	return &Store{
		NodeTypes: make(map[string]Entry),
	}
}

func (store *Store) Add(id string, factory Factory) *Store {
	store.NodeTypes[id] = Entry{
		ID:      id,
		Factory: factory,
	}
	return store
}

type Entry struct {
	ID      string
	Factory Factory
}
