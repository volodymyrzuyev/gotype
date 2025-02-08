package cursor

import (
	"github.com/volodymyrzuyev/gotype/file"
)

type Cursor interface {
	MoveJ()
	MoveK()
	MoveH()
	MoveL()
	GetRow() int
	GetCol() int
	MoveToCol(col int)
}

type cursor struct {
	row int
	col int
	f   file.File
}

func (c *cursor) updateCol() {
	if c.col >= c.f.GetLineLength(c.row) {
		c.col = max(c.f.GetLineLength(c.row)-1, 0)
	}
}

func (c *cursor) MoveJ() {
	if c.row == c.f.GetFileLength()-1 {
		return
	}
	c.row++
	c.updateCol()
}

func (c *cursor) MoveToCol(i int) {
	c.col = i
	c.updateCol()
}

func (c *cursor) MoveK() {
	if c.row == 0 {
		return
	}
	c.row--
	c.updateCol()
}

func (c *cursor) MoveH() {
	if c.col == 0 {
		return
	}
	c.col--
}

func (c *cursor) MoveL() {
	if c.col > c.f.GetLineLength(c.row)-1 {
		return
	}
	c.col++
}

func (c cursor) GetRow() int {
	return c.row
}

func (c cursor) GetCol() int {
	return c.col
}

func NewCursor(f file.File) Cursor {
	cur := cursor{
		row: 0,
		col: 0,
		f:   f,
	}
	return &cur
}
