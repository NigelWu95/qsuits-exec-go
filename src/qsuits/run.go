package qsuits

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Exec(javaPath string, jarPath string, jvmParams []string, params []string) error {

	cmd := exec.Command(javaPath)
	for i := range jvmParams {
		cmd.Args = append(cmd.Args, jvmParams[i])
	}
	cmd.Args = append(cmd.Args, "-jar")
	cmd.Args = append(cmd.Args, jarPath)
	for i := range params {
		cmd.Args = append(cmd.Args, params[i])
	}
	cmd.Stdin = os.Stdin
	// StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}
	// 创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	outReader := bufio.NewReader(stdout)
	// 实时循环读取输出流中的一行内容
	var line string
	for {
		line, err = outReader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Print(line)
	}
	errReader := bufio.NewReader(stderr)
	errRet := ""
	for {
		line, err = errReader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		errRet += line
	}
	// 阻塞直到该命令执行完成，该命令必须是被 Start 方法开始执行的
	err = cmd.Wait()
	if err != nil {
		return errors.New(strings.TrimSuffix(errRet, "\n"))
	}
	return nil
}
