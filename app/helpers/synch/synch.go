package synch

import "sync"

type Syncronizer struct {
	waitGroup *sync.WaitGroup
	rwMutex   *sync.RWMutex
	mutex     *sync.Mutex
}

// NewSyncronizer instantiates a new Syncronizer
func NewSyncronizer() *Syncronizer {
	return &Syncronizer{
		waitGroup: new(sync.WaitGroup),
		rwMutex:   new(sync.RWMutex),
		mutex:     new(sync.Mutex),
	}
}

// GetWG returns pointer of sync.WaitGroup
func (s Syncronizer) GetWG() *sync.WaitGroup {
	return s.waitGroup
}

// GetRM returns pointer of sync.RWMutex
func (s Syncronizer) GetRM() *sync.RWMutex {
	return s.rwMutex
}

// GetM returns pointer of sync.Mutex
func (s Syncronizer) GetM() *sync.Mutex {
	return s.mutex
}
