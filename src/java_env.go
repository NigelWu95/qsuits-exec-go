package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
)

func JdkDownload() {

	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", "java -version")
	fmt.Println(cmd)
	//cmd := exec.Command("java", "-version")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error(), stderr.String())
	} else {
		fmt.Println(out.String())
	}
	fmt.Println(cmd.ProcessState.Sys() == syscall.WaitStatus(0))
	fmt.Println(cmd)
}

func main()  {
	JdkDownload()
}
