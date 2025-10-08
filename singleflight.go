package singleflight

import (
	"fmt"
	"sync"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

type group struct {
	mu sync.Mutex
	m  map[string]*call
}

func (g *group) cleanup(key string, c *call) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.m[key] == c {
		delete(g.m, key)
	}
}

func (g *group) safeCall(fn func() (interface{}, error)) (val interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = toError(r)
		}
	}()
	return fn()
}

func toError(r interface{}) error {
	switch v := r.(type) {
	case error:
		return v
	default:
		return fmt.Errorf("panic: %v", v)
	}
}

func (g group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	//TODO implement me
	g.mu.Lock()
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c := &call{}
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()
	defer func() {
		g.cleanup(key, c)
	}()
	c.val, c.err = g.safeCall(fn)
	c.wg.Done()
	return c.val, c.err
}

func (g group) DoChan(key string, fn func() (interface{}, error)) <-chan Result {
	//TODO implement me
	ch := make(chan Result, 1)
	g.mu.Lock()
	if c, ok := g.m[key]; ok {
		c.chans = append(c.chans, ch)
		g.mu.Unlock()
		return ch
	}
	c := &call{chans: []chan<- Result{ch}}
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	go func() {
		defer g.cleanup(key, c)
		c.val, c.err = g.safeCall(fn)
		c.wg.Done()
		for _, ch := range c.chans {
			ch <- Result{Val: c.val, Err: c.err, Shared: len(c.chans) > 1}
		}
	}()
	return ch
}

func (g group) Forget(key string) {
	//TODO implement me
	panic("implement me")
}

type call struct {
	wg    sync.WaitGroup
	val   interface{}
	err   error
	chans []chan<- Result
}

type Result struct {
	Val    interface{}
	Err    error
	Shared bool
}

type Group interface {
	Do(key string, fn func() (interface{}, error)) (interface{}, error)
	DoChan(key string, fn func() (interface{}, error)) <-chan Result
	Forget(key string)
}

func NewGroup() Group {
	return &group{m: make(map[string]*call)}
}
