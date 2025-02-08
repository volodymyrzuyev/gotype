package view

import (
	"os"
	"strings"

	"github.com/volodymyrzuyev/gotype/cursor"
	"github.com/volodymyrzuyev/gotype/file"
	"golang.org/x/term"
)

const (
	cursorLineBG = "\033[48;5;242m"
	cursorBG     = "\033[48;5;246m"
	colorReset   = "\033[0m"
	clearScreen  = "\033[2J\033[H"
)

type View interface {
	PrintScreen(c cursor.Cursor, f file.File)
	Restore()
}

type view struct {
	fd       int
	oldState *term.State
	term     *term.Terminal
}

func (v view) Restore() {
	term.Restore(v.fd, v.oldState)
}

func (v view) PrintScreen(c cursor.Cursor, f file.File) {
	cols, rows, err := term.GetSize(v.fd)
	if err != nil {
		panic("Can't open terminal")
	}

	startingLine := max(0, c.GetRow()-10)
	endingLine := min(c.GetRow()+rows-10, f.GetFileLength())

	out := ""

	for i := startingLine; i < endingLine; i++ {
		line := padLine(f.GetLine(i), cols)
		if i == c.GetRow() {
			line = v.genCursorLine(line, c.GetCol())
		}
		out += line
	}

	v.term.Write([]byte(clearScreen))
	v.term.Write([]byte(out))
}

func (v view) genCursorLine(line string, col int) string {
	return cursorLineBG + line[:col] + cursorBG + line[col:col+1] + cursorLineBG + line[col+1:] + colorReset
}

func NewView(f file.File, c cursor.Cursor) View {
	v := view{}

	v.fd = int(os.Stdout.Fd())
	oldState, err := term.MakeRaw(v.fd)

	if err != nil {
		panic("Error Creating Terminal")
	}

	v.oldState = oldState
	v.term = term.NewTerminal(os.Stdout, " ")

	return v
}

func padLine(line string, cols int) string {
	nl := strings.ReplaceAll(line, "	", "    ")
	return nl + strings.Repeat(" ", cols-len(nl))
}
