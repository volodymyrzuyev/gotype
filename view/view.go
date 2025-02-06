package view

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/volodymyrzuyev/gotype/cursor"
	"github.com/volodymyrzuyev/gotype/file"
)

const (
	cursorLineColorForground  = "\033[38;5;232m"
	cursorLineColorBackground = "\033[48;5;255m"
	cursorColorBackground     = "\033[48;5;251m"
	colorReset                = "\033[0m"
	clearScreen               = "\033[2J\033[H"
)

var (
	cursorLineBG = "\033[48;5;242m"
	cursorBG     = "\033[48;5;246m"
)

type View interface {
	PrintScreen(c cursor.Cursor, f file.File)
}

type view struct {
}

func getTerminalDimensions() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	size := strings.Fields(string(output))

	rows, _ := strconv.Atoi(size[0])
	cols, _ := strconv.Atoi(size[1])
	fmt.Println(rows, cols)
	return rows, cols
}

func (v view) PrintScreen(c cursor.Cursor, f file.File) {
	rows, cols := getTerminalDimensions()

	startingLine := max(0, c.GetRow()-10)
	endingLine := min(c.GetRow()+rows-10, f.GetFileLength())

	out := ""

	for i := startingLine; i < endingLine; i++ {
		line := f.GetLine(i)
		pLine := line + strings.Repeat(" ", cols-len(line))
		if i == c.GetRow() {
			col := c.GetCol()
			pLine = cursorLineBG + pLine[:col] + cursorBG + pLine[col:col+1] + cursorLineBG + pLine[col+1:] + colorReset
		}
		out += pLine + "\n"
	}

	fmt.Print(clearScreen)
	fmt.Print(out)
}

func NewView(f file.File, c cursor.Cursor) View {
	v := view{}

	return v
}
