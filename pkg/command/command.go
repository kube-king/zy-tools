/*
Copyright 2022 qkp Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package command

import (
	"context"
	"errors"
	"io"
	"os/exec"
	"strings"
	"time"
)

type Exec struct {
	Command string
	Args    []string
	Timeout time.Duration
	Stdout  io.Writer
	Stderr  io.Writer
	Cmd     *exec.Cmd
	Dir     string
}

type ExecResponse struct {
	ExitCode int
	Success  bool
}

func NewExec(command string, args ...string) *Exec {
	return &Exec{
		Command: command,
		Args:    args,
	}
}

// SetTimeout 设置命令执行超时时间
func (e *Exec) SetTimeout(duration time.Duration) *Exec {
	e.Timeout = duration
	return e
}

func (e *Exec) SetDir(dir string) *Exec {
	e.Dir = dir
	return e
}

// SetArgs 设置命令执行参数
func (e *Exec) SetArgs(args ...string) *Exec {
	e.Args = args
	return e
}

// SetStdout 设置命令执行返回正常消息
func (e *Exec) SetStdout(stdout io.Writer) *Exec {
	e.Stdout = stdout
	return e
}

// SetStderr 设置命令执行返回错误消息
func (e *Exec) SetStderr(stderr io.Writer) *Exec {
	e.Stderr = stderr
	return e
}

// SetCommand 设置执行命令
func (e *Exec) SetCommand(command string) *Exec {
	e.Command = command
	return e
}

// Run 执行命令
func (e *Exec) Run() (response *ExecResponse, err error) {
	var cmd *exec.Cmd
	var ctx context.Context
	var cancel context.CancelFunc
	response = &ExecResponse{}

	if e.Timeout <= 0 {
		cmd = exec.CommandContext(context.Background(), e.Command, e.Args...)
	} else {
		context.WithTimeout(context.Background(), e.Timeout)
		ctx, cancel = context.WithTimeout(context.Background(), e.Timeout)
		defer cancel()
		cmd = exec.CommandContext(ctx, e.Command, e.Args...)
	}
	if e.Dir != "" {
		cmd.Dir = e.Dir
	}
	cmd.Stdout = e.Stdout
	cmd.Stderr = e.Stderr

	args := append([]string{e.Command}, e.Args...)
	cmd.Stdout.Write([]byte(strings.Join(args, " ")))

	err = cmd.Run()
	response.ExitCode = cmd.ProcessState.ExitCode()

	if ctx != nil && ctx.Err() != nil && errors.Is(ctx.Err(), context.DeadlineExceeded) {
		err = errors.New("Exec Running Timeout ")
		response.Success = false
		return
	}
	if err != nil {
		return
	}

	return response, nil
}
