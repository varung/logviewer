package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWidth() uint {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return uint(ws.Col)
}

func ColWidth() uint {
	return 120
}

func main() {
	log.SetFlags(log.Lshortfile)
	fns := []*bufio.Reader{}
	for _, f := range os.Args[1:] {
		file, err := os.Open(f)
		if err != nil {
			log.Panic(err)
		}
		fns = append(fns, bufio.NewReader(file))
	}

	numfiles := len(fns)
	colwidth := int(ColWidth()) / numfiles
	colspace := strings.Repeat(" ", colwidth)

	// read from each file, parsing timestamp, and then output it appropriately
	lines := []string{}
	for i := 0; i < numfiles; i++ {
		l, _ := fns[i].ReadString('\n')
		lines = append(lines, l)
		name := os.Args[1+i]
		fmt.Print(name)
		fmt.Print(strings.Repeat(" ", colwidth-len(name)))
	}
	fmt.Println()

	// print out table headings

	for {
		minline := 0
		for i := 1; i < numfiles; i++ {
			l := lines[i]
			//log.Print(minline, i, l)
			if lines[minline] == "" ||
				(l != "" && (l < lines[minline])) {
				minline = i
			}
		}
		if lines[minline] == "" {
			break
		}
		fmt.Print(strings.Repeat(colspace, minline) + lines[minline])
		l, _ := fns[minline].ReadString('\n')
		lines[minline] = l
	}
}
