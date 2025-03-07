package clients

import (
	"context"
	"os/exec"
)

type CmdClient struct{}

func NewCmdClient() *CmdClient {
	return &CmdClient{}
}

func (c *CmdClient) Command(ctx context.Context, cmd string, args []string) (string, error) {
	e := exec.CommandContext(ctx, cmd, args...)
	stdout, err := e.Output()
	//handle err a little better - pg_ctl will return a uninstantiaed
	//database directory as an error
	if err != nil && len(stdout) == 0 {
		return "", err
	}

	return string(stdout), nil
}
