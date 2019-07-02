package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

var cmdWithPath string
var err error
var cmd = "ls -l"

func Init() {
	cmdWithPath, err = exec.LookPath("bash")
	if err != nil {
		fmt.Println("not find bash.")
		os.Exit(5)
	}
}

func method1() {
	cmd := exec.Command(cmdWithPath, "-c", cmd)
	//cmd := exec.Command("java", "-version")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
	fmt.Println(cmd.ProcessState.Sys() == syscall.WaitStatus(0))
	fmt.Println(cmd.Stdin)
	fmt.Println(cmd.Stdout)
	fmt.Println(cmd)
}

func method2()  {
	//cmd := exec.Command("java", "-version")
	cmd := exec.Command(cmdWithPath, "-c", cmd)
	//cmd := exec.Command(cmdWithPath, "-c", "java -version")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf(string(output))
	fmt.Println(cmd.ProcessState.Sys() == syscall.WaitStatus(0))
	//fmt.Println(cmd)
}

func method3()  {
	cmd := exec.Command(cmdWithPath, "-c", cmd)
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil{
		fmt.Println("Execute failed when Start:" + err.Error())
		return
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return
	}
	stdin.Write([]byte("go text for grep\n"))
	stdin.Write([]byte("go test text for grep\n"))
	stdin.Close()
	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()
	fmt.Println("Execute finished:" + string(out_bytes))
	//fmt.Println(cmd)
}

func main()  {
	Init()
	method1()
	method2()
	method3()
}