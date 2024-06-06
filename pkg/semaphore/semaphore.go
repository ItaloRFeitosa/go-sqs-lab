package semaphore

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	ch chan struct{}
}

func New(maxConcurreny int) Semaphore {
	return &semaphore{
		ch: make(chan struct{}, maxConcurreny),
	}
}

func (s *semaphore) Acquire() {
	s.ch <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.ch
}
