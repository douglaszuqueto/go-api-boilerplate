package graceful

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Signals Signals
type Signals struct {
	signalCh chan os.Signal
	doneCh   chan bool
}

// New New
func New() *Signals {
	s := &Signals{
		signalCh: make(chan os.Signal, 1),
		doneCh:   make(chan bool, 1),
	}

	signal.Notify(s.signalCh, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-s.signalCh
		log.Println("Signal stop:", sig)

		s.doneCh <- true
	}()

	return s
}

// Wait Wait
func (s *Signals) Wait() {
	<-s.doneCh
}
