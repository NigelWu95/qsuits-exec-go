package main

import (
	"fmt"
	"os"
	"os/exec"
)

var cmdWithPath string
var err error

func Init() {
	cmdWithPath, err = exec.LookPath("bash")
	if err != nil {
		fmt.Println("not find bash.")
		os.Exit(5)
	}
}

func method1() {
	//cmd := exec.Command(cmdWithPath, "-c", "ls")
	cmd := exec.Command("java", "-version")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
}

func method2()  {
	//cmd := exec.Command("java", "-version")
	//cmd := exec.Command(cmdWithPath, "-c", "ls -l")
	cmd := exec.Command(cmdWithPath, "-c", "java -version")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf(string(output))
}

func main()  {
	Init()
	method1()
	method2()
}