package main

import (
	"Go-algorithm/src/links"
	"fmt"
	"strconv"
)

func main() {
	head := links.InitCircleLink(10)
	head.Add("aaaa")
	head.Add("bbbb")
	head.Add("cccc")
	head.Add("dddd")
	if err := head.RemoveOne(0); err != nil {
          fmt.Println(err)
	}
	if err := head.RemoveOne(3); err != nil {
		fmt.Println(err)
	}
	for i:=0;i<10;i++ {
		head.Add(strconv.Itoa(i))
	}
	for j:=3;j<8 ;j++{
		if err := head.RemoveOne(j); err != nil {
			fmt.Println(err)
		}
	}
	for i:=10;i<18;i++ {
		head.Add(strconv.Itoa(i))
	}
	head.ShowList()
}
