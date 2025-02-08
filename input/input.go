package input

import (
	"os"

	"golang.org/x/term"
)

type Input interface {
	CatchInputs()
	Restore()
}

type input struct {
	oldState *term.State
	fd       int
	inp      chan byte
}

func (i input) CatchInputs() {
	buf := make([]byte, 1)
	for {
		os.Stdin.Read(buf)
		if len(buf) < 1 {
			continue
		}
		i.inp <- buf[0]
	}
}

func (i input) Restore() {
	term.Restore(i.fd, i.oldState)
}

func NewInput(ch chan byte) Input {
	fd := int(os.Stdin.Fd())
	t, err := term.MakeRaw(fd)
	if err != nil {
		panic("Can not capture STDin")
	}

	i := input{
		oldState: t,
		fd:       fd,
		inp:      ch,
	}

	return &i
}
