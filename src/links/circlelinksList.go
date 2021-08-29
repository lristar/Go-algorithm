package links

import (
	"errors"
	"fmt"
)

//循环单链表
type data string

type CircleHead struct {
	Head *CircleNode
	len  int
	size int
}

type CircleNode struct {
	index int
	Value data
	Next  *CircleNode
}

func InitCircleLink(size int) *CircleHead {
	return &CircleHead{
		size: size,
	}
}

func (c *CircleHead) Add(str string) {
	da := &CircleNode{
		Value: data(str),
	}
	if !c.IsFull() {
		if last := c.GetLast(); last != nil {
			da.index = last.index + 1
			last.Next = da
			da.Next = c.Head
			c.len++
		} else {
			c.Head = da
			da.Next = da
		}
	}
	return

}

func (c *CircleHead) RemoveAll() {
	c.Head = nil
	c.len = 0
}

func (c *CircleHead) RemoveOne(index int) error {
	if !c.IsEmpty() {
		//取第一个节点
		head := c.Head
		//删除头结点并且节点为一
		if index == head.index {
			if c.len == 1 {
				c.RemoveAll()
				return nil
			}
			last := c.GetLast()
			c.Head = head.Next
			last.Next = c.Head
			c.len--
			return nil

		}
		//头结点和尾结点之间的节点
		for head.Next != c.Head {
			if head.Next.index == index {
				head.Next = head.Next.Next
				c.len--
				return nil
			}
			head = head.Next
		}
		//与尾结点相同
		if head.Next.index == index {
			head.Next = c.Head
			c.len--
			return nil
		}
		return errors.New("no find")
	}
	return errors.New("link is empty")
}

func (c *CircleHead) ShowList() {
	if c.IsEmpty() {
		fmt.Println("link is empty")
		return
	}
	data := c.Head
	fmt.Println("head:", data)
	for data.Next != c.Head {
		data = data.Next
		fmt.Println("data:", data)
	}
}

func (c *CircleHead) IsEmpty() bool {
	if c.Head == nil {
		return true
	}
	return false
}

func (c *CircleHead) IsFull() bool {
	if c.GetLength() == c.GetSize()-1 {
		return true
	}
	return false
}

func (c *CircleHead) GetSize() int {
	return c.size
}

func (c *CircleHead) GetLength() int {
	return c.len
}

func (c *CircleHead) GetHead() (int, data) {
	return c.Head.index, c.Head.Value
}

func (c *CircleHead) GetLast() *CircleNode {
	if c.IsEmpty() {
		return nil
	}
	data := c.Head
	for data.Next != c.Head {
		data = data.Next
	}
	return data
}
