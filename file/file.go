package file

import (
	"bufio"

	"github.com/volodymyrzuyev/gotype/logger"
)

type File interface {
	InsertLine(newLine string, afterLine int)
	ChangeLine(newLine string, line int)
	DeleteLine(line int)
	GetLineLength(line int) int
	GetLine(line int) string
	GetFileLength() int
	IsLineEmpty(line int) bool
}

type file struct {
	lines     []string
	fileLines int // 0 based
}

func (f *file) InsertLine(newLine string, afterLine int) {
	logger.INFO.Printf("File Pre Enter new line:\n %v", f.getWholeFile())
	newLines := make([]string, f.fileLines+1)
	newLines[afterLine+1] = newLine

	for i, l := range f.lines {
		ni := i

		if i > afterLine {
			ni += 1
		}

		newLines[ni] = l

	}

	f.lines = newLines

	f.fileLines += 1
	logger.INFO.Printf("File Post Enter new line:\n %v", f.getWholeFile())
}

func (f *file) ChangeLine(newLine string, line int) {
	if line > f.fileLines {
		return
	}

	f.lines[line] = newLine
}

func (f *file) DeleteLine(line int) {
	if line <= 0 {
		return
	}

	newLines := make([]string, f.fileLines)
	for i, l := range f.lines {
		j := i
		if i == line {
			continue
		}
		if i > line {
			j = i - 1
		}
		newLines[j] = l
	}

	f.lines = newLines
}

func (f file) GetFileLength() int {
	return f.fileLines
}

func (f file) GetLineLength(line int) int {
	return len(f.lines[line])
}

func (f file) GetLine(line int) string {
	return f.lines[line]
}

func (f file) IsLineEmpty(line int) bool {
	return f.lines[line] == ""
}

func (f file) getWholeFile() string {
	out := ""
	for i := range f.lines {
		out += f.lines[i] + "\n"
	}

	return out
}

func NewFile(fileScaner *bufio.Scanner) File {
	var lines []string
	for fileScaner.Scan() {
		lines = append(lines, fileScaner.Text())
	}

	newF := file{
		lines:     lines,
		fileLines: len(lines),
	}

	return &newF
}
