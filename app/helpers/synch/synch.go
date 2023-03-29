package synch

import (
	"sync"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/chanh"
)

type Syncronizer struct {
	waitGroup *sync.WaitGroup
	rwMutex   *sync.RWMutex
	mutex     *sync.Mutex
	errChan   chan error
}

// NewSyncronizer instantiates a new Syncronizer
func NewSyncronizer() *Syncronizer {
	return &Syncronizer{
		waitGroup: new(sync.WaitGroup),
		rwMutex:   new(sync.RWMutex),
		mutex:     new(sync.Mutex),
		errChan:   chanh.Make[error](nil, 1),
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

// Err returns value of errChan, if argument is provided it will set "errChan" with the given argument
func (s Syncronizer) Err(errs ...error) error {
	if len(errs) >= 1 {
		err := errs[0]
		chanh.ReplaceVal(s.errChan, err)
	}
	return chanh.GetValAndKeep(s.errChan)
}
