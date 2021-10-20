package main

import (
	"os/exec"
	"time"
)

func main() {
	for true{
		exec.Command("/bin/bash","-c" ,"echo 'aaaa' >> /home/lzy/test/aaaaaa")
		time.Sleep(600)
	}
}
