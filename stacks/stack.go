package stacks

import (
	"errors"
)

type Stack[T any] struct {
	data    []T
	current int
}

func Test() {

}
func (ds *Stack[T]) Peek() (T, error) {
	var item T
	if ds.current >= 0 {
		item = ds.data[ds.current]
		return item, nil
	}
	return item, errors.New("empty stack. cant peek")
}

func (ds *Stack[T]) Push(item T) error {
	ds.data = append(ds.data, item)
	ds.current++
	return nil
}

func (ds *Stack[T]) Pop() (T, error) {
	var item T
	if ds.current >= 0 {
		item = ds.data[ds.current]
		ds.current--
		ds.data = ds.data[:ds.current]
		return item, nil
	}
	return item, errors.New("empty stack. cant pop")
}
