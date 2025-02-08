package main

import (
	"bufio"
	"os"
	"time"

	"github.com/volodymyrzuyev/gotype/cursor"
	"github.com/volodymyrzuyev/gotype/file"
	"github.com/volodymyrzuyev/gotype/input"
	"github.com/volodymyrzuyev/gotype/logger"
	"github.com/volodymyrzuyev/gotype/view"
)

func openFile() file.File {
	d, _ := os.Open("test.txt")
	k := bufio.NewScanner(d)
	f := file.NewFile(k)
	d.Close()
	return f
}

func main() {
	logger.InitLogger("goType.log")
	logger.INFO.Println("Starting goType")

	f := openFile()
	c := cursor.NewCursor(f)

	v := view.NewView(f, c)
	ch := make(chan byte)
	i := input.NewInput(ch)

	d := cursor.NewDecoder(c, f)

	go i.CatchInputs()

	logger.INFO.Println("Init successful")

	for {
		select {
		case in, _ := <-ch:
			switch in {
			case 'q':
				i.Restore()
				v.Restore()
				logger.CloseLogger()
				os.Exit(0)
			case 'w':
				saveFile(f)
			default:
				d.ParseInput(in)
			}
		default:
			v.PrintScreen(c, f)
		}
		time.Sleep(time.Second / 60)
	}
}

func saveFile(f file.File) {
	var bates []byte
	for i := range f.GetFileLength() {
		bates = append(bates, []byte(f.GetLine(i)+"\n")...)
	}
	os.WriteFile("test.txt", bates, 0644)
}
