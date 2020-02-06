package stack

import "errors"

type ArrayStack struct {
	elements []interface{}
	size int
}

func assertArrayStackImplementation()  {
	var _ Stack = (*ArrayStack)(nil)
}

func NewDefault() *ArrayStack {
	return &ArrayStack{
		elements: make([]interface{}, 10),
		size:     0,
	}
}

func NewArrayStack(capacity int) *ArrayStack {
	return &ArrayStack{
		elements: make([]interface{}, capacity),
		size: 0,
	}
}

func (s *ArrayStack) GetSize() int{
	return s.size
}

func (s *ArrayStack) Pop() (interface{}, error){
	if s.size == 0 {
		return nil, errors.New("栈空间为零")
	}
	s.size --
	return s.elements[s.size], nil
}

func (s *ArrayStack) Peek() (interface{}, error) {
	if s.size == 0 {
		return nil, errors.New("栈空间为零")
	}
	s.size --
	return s.elements[s.size], nil
}

func (s *ArrayStack) Push(e interface{})  error {
	if s.size == cap(s.elements) {
		return errors.New("栈空间已满，无法再存入元素。")
	}
	s.elements[s.size] = e
	s.size ++
	return nil
}

func (s *ArrayStack) Resize() error {
	return nil
}
