package cursor

import (
	"github.com/volodymyrzuyev/gotype/file"
	"github.com/volodymyrzuyev/gotype/logger"
)

type mode int

const (
	scrolling mode = 0
	typing    mode = 1
)

type decoder struct {
	m mode
	f file.File
	c Cursor
}

func (d *decoder) ParseInput(in byte) int {
	if d.m == typing {
		d.decodeTyping(in)
		return 0
	}

	if d.m == scrolling {
		if in == 'q' {
			return 10
		}

		if in == 'w' {
			return 11
		}

		d.decodeScrolling(in)
		return 0
	}

	return 0
}

func (d *decoder) decodeTyping(in byte) {
	switch in {
	case byte(27):
		d.m = scrolling
		break
	case byte(13):
		logger.INFO.Println("Enter Pressed")
		d.newLine("")
		break
	case byte(127):
		d.backSpace()
		break
	default:
		d.updateLine(string(in))
	}
}

func (d decoder) backSpace() {
	row := d.c.GetRow()
	if d.f.IsLineEmpty(row) {
		d.f.DeleteLine(row)
		d.c.MoveK()
		return
	}
	cLine := d.f.GetLine(row)

	col := d.c.GetCol()
	if col == 0 {
		pLine := d.f.GetLine(row - 1)

		d.f.ChangeLine(pLine+cLine, row-1)
		d.c.MoveK()
		d.c.MoveToCol(len(pLine))
		d.f.DeleteLine(row)
		return
	}

	nLine := ""

	if len(cLine) > 1 {
		nLine = cLine[:col-1] + cLine[col:]
	}

	d.f.ChangeLine(nLine, row)
	d.c.MoveH()
}

func (d decoder) newLine(line string) {
	row := d.c.GetRow()

	if !d.f.IsLineEmpty(row) {
		col := d.c.GetCol()

		cLine := d.f.GetLine(row)
		d.f.ChangeLine(cLine[:col], row)
		line += cLine[col:]
	}
	d.f.InsertLine(line, row)
	d.c.MoveJ()
	d.c.MoveToCol(0)
}

func (d decoder) updateLine(char string) {
	row := d.c.GetRow()
	col := d.c.GetCol()

	cLine := d.f.GetLine(row)
	nLine := cLine[:col] + char + cLine[col:]
	d.f.ChangeLine(nLine, row)

	d.c.MoveL()
}

func (d *decoder) decodeScrolling(in byte) {
	switch in {
	case 'j':
		d.c.MoveJ()
		break
	case 'k':
		d.c.MoveK()
		break
	case 'h':
		d.c.MoveH()
		break
	case 'l':
		d.c.MoveL()
		break
	case 'a':
		d.c.MoveL()
		d.m = typing
		break
	case 'i':
		d.m = typing
		break
	}
}

func NewDecoder(c Cursor, f file.File) decoder {
	d := decoder{
		m: scrolling,
		f: f,
		c: c,
	}

	return d
}
