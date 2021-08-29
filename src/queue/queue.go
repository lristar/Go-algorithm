package queue

import "errors"

//循环队列
type Queue struct {
	List []string
	Tail int
	Head int
	Size int
}

func NewQueue(size int)*Queue{
	list:=make([]string,size)
	return &Queue{
		List: list,
		Size: size,
	}
}

func (q *Queue) Push(str string) {
	if !q.IsFull() {
		q.List[q.Head]=str
		q.Head = (q.Head + 1) % q.Size
	}
}

func (q *Queue) Pop()(string,error) {
	if !q.IsEmpty() {
		str:=q.List[q.Tail]
		q.Tail = (q.Tail + 1) % q.Size
		return str,nil
	}
	return "",errors.New("queue is empty")
}

func (q *Queue) IsFull() bool {
	if (q.Head-q.Tail+q.Size)%q.Size == q.Size-1 {
		return true
	}
	return false
}

func (q *Queue) IsEmpty() bool {
	if q.Head == q.Tail {
		return true
	}
	return false
}

func (q *Queue) ShowList() []string {
	var list []string
	if q.Tail > q.Head {
		list = append(list, q.List[q.Tail:q.Size]...)
		list = append(list,q.List[0:q.Head]...)
		return list
	}
	return q.List[q.Tail:q.Head]
}
