package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var ErrPoolClosed = errors.New("pool is closed")

type Pool struct {
	factory   func() (io.Closer, error)
	resources chan io.Closer
	closed    bool
	mutex     *sync.Mutex
}

func New(poolSize int, factory func() (io.Closer, error)) (*Pool, error) {
	p := &Pool{
		factory:   factory,
		resources: make(chan io.Closer, poolSize),
		closed:    false,
		mutex:     &sync.Mutex{},
	}

	return p, nil
}

//methods of Pool type
func (p *Pool) Acquire() (io.Closer, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	select {
	case r, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		fmt.Println("Acquiring from the pool")
		return r, nil
	default:
		fmt.Println("Acquiring from the factory")
		return p.factory()
	}
}

func (p *Pool) Release(resource io.Closer) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	select {
	case p.resources <- resource:
		fmt.Println("Releasing the resource to the pool")
		return nil
	default:
		fmt.Println("Pool full. Discarding the resource")
		return resource.Close()
	}
}

func (p *Pool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.resources)
	for resource := range p.resources {
		resource.Close()
	}
}

/*
Aquire() => resource
Release(resource)
Close()
*/
