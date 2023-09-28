package code

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
	"time"
)

type service struct {
}

func NewService() service {
	return service{}
}

func (s *service) Run(ctx context.Context, code Code) (Output, error) {
	var output Output
	ctx, cancel := context.WithTimeout(
		ctx,
		1*time.Second,
	)
	cmd := exec.CommandContext(ctx, "./golox")

	var outb, errb bytes.Buffer

	cmd.Stdout = &outb
	cmd.Stderr = &errb
	cmd.Stdin = strings.NewReader(code.Text)

	go func() {
		cmd.Run()
		cancel()
	}()

	<-ctx.Done()
	switch ctx.Err() {
	case context.DeadlineExceeded:
		output.Message = "Time limit exceeded."
	case context.Canceled:
		output.Output = outb.String()
	default:
		return Output{}, ctx.Err()
	}

	return output, nil
}
