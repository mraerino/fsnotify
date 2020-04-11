// +build js

package fsnotify

import "errors"

type Watcher struct {
	Events chan Event
	Errors chan error
}

func NewWatcher() (*Watcher, error) {
	return &Watcher{
		Events: make(chan Event),
		Errors: make(chan error),
	}, nil
}

func (w *Watcher) Close() error {
	close(w.Events)
	close(w.Errors)
	return nil
}

var ErrNotImplemented = errors.New("not implemented")

func (*Watcher) Add(name string) error {
	return ErrNotImplemented
}

func (*Watcher) Remove(name string) error {
	return ErrNotImplemented
}
