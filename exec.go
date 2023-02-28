package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

type ExecResult struct {
	Stdout string
	Stderr string
	Cmderr string
}

/*
kill geth 进程,由于当前docker 在geth进程被kill后会自动重启
,RestartDocker可能在未获取到执行结果时,宿主docker就被重启了
此时返回结果不具备实际意义
*/
func RestartDocker(name string) ExecResult {
	sendMsg(fmt.Sprintf("docker:%s restart", name))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", fmt.Sprintf("docker restart %s ", name))
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	Logger.Sugar().Info("kill ", name)
	err := cmd.Run()

	if stderr.Bytes() != nil {
		Logger.Sugar().Error(stderr.String())
	}
	var errStr = ""
	if err != nil {
		Logger.Sugar().Error(err.Error())
		errStr = err.Error()
	}

	res := ExecResult{
		stdout.String(),
		stderr.String(),
		errStr,
	}
	return res
}
