package stack

import (
	"errors"
	"strings"
)

type Stack struct {
	Top int
	Str []string
}

func (s *Stack) Push(str string) {
	s.Str = append(s.Str, str)
	s.Top++
}

func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() == 1 {
		s.Top--
		str:=s.Str[s.Top]
		s.Str=s.Str[0:s.Top]
		return str, nil
	}
	return "", errors.New("is empty")
}

//判断是否为空
func (s *Stack) IsEmpty() int {
	if s.Top == 0 {
		return 0
	}
	return 1
}
func (s *Stack) ToString() string {
	return strings.Join(s.Str, ",")
}


