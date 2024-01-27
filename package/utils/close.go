package utils

type Closer chan struct{}

func NewCloser() Closer { return make(Closer) }
func (c Closer) Wait()  { <-c }
func (c Closer) Close() { close(c) }
