package links

import (
	"errors"
	"fmt"
)

//单链表
type ListNode struct {
	index int
	Value string
	Next  *ListNode
}
type HeadNode struct {
	Next  *ListNode
}

func (h *HeadNode) Push(str string) {
	a := &ListNode{Value: str}
	if h.Next!=nil{
		a.index = h.Next.index + 1
	}else{
		a.index++
	}
	a.Next=h.Next
	h.Next=a
}

//刪除指定index的元素
func (h *HeadNode) Pop(index int) (string, error) {
	str := ""
	if !h.IsEmpty() {
		temp := h.Next
		if index < 0 || index > h.GetLength() {
			return "", errors.New("index有誤")
		}
		//第一個的時候直接判斷
		if temp.index==index{
			str=h.Next.Value
			h.Next=h.Next.Next
			return str,nil
		}
		//第二個后往後找
		for temp.Next != nil {
			if index == temp.Next.index {
				str = temp.Value
				temp.Next = temp.Next.Next
				return str, nil
			}
			temp = temp.Next
			fmt.Println(2)
		}

		return "", errors.New("no find")
	}
	return "", errors.New("link is empty")
}

//判斷列表是否為空
func (h *HeadNode) IsEmpty() bool {
	if h.Next == nil {
		return true
	}
	return false
}

// 獲取列表長度
func (h *HeadNode) GetLength() int {
	length := 0
	data := h.Next
	for data != nil {
		length++
		data=data.Next
	}
	return length
}

func (h *HeadNode) ShowList() {
	if h.IsEmpty() {
		return
	}
	temp := h.Next
	for temp.Next != nil {
		fmt.Println(temp.Value, ",",temp.index)
		temp = temp.Next
	}
	fmt.Println(temp.Value, ",",temp.index)
}
