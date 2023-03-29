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

// WG returns pointer of sync.WaitGroup
func (s Syncronizer) WG() *sync.WaitGroup {
	return s.waitGroup
}

// RM returns pointer of sync.RWMutex
func (s Syncronizer) RM() *sync.RWMutex {
	return s.rwMutex
}

// M returns pointer of sync.Mutex
func (s Syncronizer) M() *sync.Mutex {
	return s.mutex
}
