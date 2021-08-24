package chain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChain(t *testing.T) {
	foo := NewFoo()
	s1 := foo.Foo("foo")
	assert.Equal(t, "foo & f1 & f2 & f3", s1)
	s2 := foo.Bar("bar")
	assert.Equal(t, "bar & b1 & b2", s2)
}

func NewFoo() FooInterface {
	return New(
		func(next interface{}) interface{} {
			return &f1{FooInterface: next.(FooInterface)}
		},
		func(next interface{}) interface{} {
			return &f2{FooInterface: next.(FooInterface)}
		},
		func(next interface{}) interface{} {
			return &f3{FooInterface: next.(FooInterface)}
		},
	)(&f{}).(FooInterface)
}

type FooInterface interface {
	Foo(s string) string
	Bar(s string) string
}

type f struct {
}

func (f) Foo(s string) string {
	return s
}

func (f) Bar(s string) string {
	return s
}

type f1 struct {
	FooInterface
}

func (r f1) Foo(s string) string {
	return r.FooInterface.Foo(fmt.Sprintf("%s & %s", s, "f1"))
}

func (r f1) Bar(s string) string {
	return r.FooInterface.Bar(fmt.Sprintf("%s & %s", s, "b1"))
}

type f2 struct {
	FooInterface
}

func (r f2) Foo(s string) string {
	return r.FooInterface.Foo(fmt.Sprintf("%s & %s", s, "f2"))
}

func (r f2) Bar(s string) string {
	return fmt.Sprintf("%s & %s", s, "b2")
}

type f3 struct {
	FooInterface
}

func (r f3) Foo(s string) string {
	return r.FooInterface.Foo(fmt.Sprintf("%s & %s", s, "f3"))
}

func (r f3) Bar(s string) string {
	return r.FooInterface.Bar(fmt.Sprintf("%s & %s", s, "b3"))
}
