package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func JdkDownload() {

	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	//cmd := exec.Command("/bin/bash", "-c", "java -version")
	cmd := exec.Command("java", "-version")
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
	//fmt.Println(cmd.ProcessState.Sys() == syscall.WaitStatus(0))
	result := fmt.Sprintln(cmd)
	fmt.Println(result)
	//fmt.Println(strings.Split(result, "  <nil>  ")[1])

	fmt.Println("Path: " + cmd.Path)
	fmt.Print("Args: ")
	fmt.Println(cmd.Args)
	fmt.Print("Env: ")
	fmt.Println(cmd.Env)
	fmt.Println("Dir: " + cmd.Dir)
	fmt.Print("SysProcAttr: ")
	fmt.Println(cmd.SysProcAttr)
	fmt.Print("ExtraFiles: ")
	fmt.Println(cmd.ExtraFiles)
	fmt.Println(cmd.Process)
	fmt.Println(stderr.String())
}

func main()  {
	JdkDownload()
}
