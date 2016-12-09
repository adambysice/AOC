package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
)

type vec struct{
	x,
	y int
}

func main() {
	var code string
	reader := bufio.NewReader(os.Stdin)
	for {
		pos := vec{0,0}
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("eof")
			break
		}
		for i := range text {
			s := string(text[i])
			if s == "U" {
				pos.y++
			} else if s =="R" {
				pos.x++
			} else if s == "D" {
				pos.y--
			} else if s == "L" {
				pos.x--
			}
			//clamp
			if pos.x > 1 {
				pos.x = 1
			}
			if pos.y > 1 {
				pos.y = 1
			}
			if pos.x < -1 {
				pos.x = -1
			}
			if pos.y < -1 {
				pos.y = -1
			}

		}
		switch pos {
			case vec{-1, 1}:
				code += "1"
			case vec{0,1}:
				code += "2"
			case vec{1,1}:
				code += "3"
			case vec{-1,0}:
				code += "4"
			case vec{0,0}:
				code += "5"
			case vec{1,0}:
				code += "6"
			case vec{-1, -1}:
				code += "7"
			case vec{0,-1}:
				code += "8"
			case vec{1,-1}:
				code += "9"
		}

	}
	fmt.Println(code)
}
