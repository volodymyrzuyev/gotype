package main

import (
	"bufio"
	"os"
	"os/exec"
	"time"

	"github.com/volodymyrzuyev/gotype/cursor"
	"github.com/volodymyrzuyev/gotype/file"
	"github.com/volodymyrzuyev/gotype/view"
)

func captureKeys(ch chan string) {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		ch <- string(b)
	}
}

func openFile() file.File {
	d, _ := os.Open("test.txt")
	k := bufio.NewScanner(d)
	f := file.NewFile(k)
	d.Close()
	return f
}

func main() {
	f := openFile()
	c := cursor.NewCursor(f)

	v := view.NewView(f, c)
	ch := make(chan string)
	go captureKeys(ch)

	mode := 0

	for {
		select {
		case stdin, _ := <-ch:
			if mode == 0 {
				getKeyPress(stdin, c, &mode, &f)
			} else {
				typeBuffer(stdin, c, &mode, f)
			}
		default:
			v.PrintScreen(c, f)
		}
		time.Sleep(time.Second / 60)
	}
}

func getKeyPress(in string, c cursor.Cursor, mode *int, f *file.File) {
	switch in {
	case "j":
		c.MoveJ()
		break
	case "k":
		c.MoveK()
		break
	case "h":
		c.MoveH()
		break
	case "l":
		c.MoveL()
		break
	case "i":
		*mode = 1
		break
	case "a":
		c.MoveL()
		*mode = 1
		break
	case "w":
		saveFile(*f)
		*f = openFile()
		break

	}
}

func saveFile(f file.File) {
	var bates []byte
	for i := range f.GetFileLength() {
		bates = append(bates, []byte(f.GetLine(i)+"\n")...)
	}
	os.WriteFile("test.txt", bates, 0644)
}

func typeBuffer(in string, c cursor.Cursor, mode *int, f file.File) {
	switch in {
	case string(rune(27)):
		*mode = 0
		break
	case string(rune(10)):
		addLine(c, f)
		break
	case string(rune(127)):
		backSpace(c, f)
	default:
		updateLine(in, c, f)
	}
}

func updateLine(in string, c cursor.Cursor, f file.File) {
	curLine := f.GetLine(c.GetRow())
	newLine := curLine[:c.GetCol()] + in + curLine[c.GetCol():]
	f.ChangeLine(newLine, c.GetRow())
	c.MoveL()
}

func backSpace(c cursor.Cursor, f file.File) {
	row := c.GetRow()
	if f.IsLineEmpty(row) {
		f.DeleteLine(row)
		c.MoveK()
		return
	}

	cL := f.GetLine(row)
	if c.GetCol() == 0 {
		pL := f.GetLine(row - 1)
		f.ChangeLine(pL+cL, row-1)
		f.DeleteLine(row)
		c.MoveK()
		return
	}
	nL := ""
	if len(cL) > 1 {
		nL = cL[:c.GetCol()-1] + cL[c.GetCol():]
	}
	f.ChangeLine(nL, row)
	c.MoveH()
}

func addLine(c cursor.Cursor, f file.File) {
	newL := ""

	if !f.IsLineEmpty(c.GetRow()) {
		cLine := f.GetLine(c.GetRow())
		lLine := cLine[:c.GetCol()]
		f.ChangeLine(lLine, c.GetRow())
		newL = cLine[c.GetCol():]
	}

	f.InsertLine(newL, c.GetRow())
	c.MoveJ()
	c.MoveH()
}
