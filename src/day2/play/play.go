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

var keyboard map[vec]string = map[vec]string{
	vec{2,0}: "1",
	vec{-1,1}: "2",
	vec{0, 1}: "3",
	vec{1,1}: "4",
	vec{-2,0}: "5",
	vec{-1,0}: "6",
	vec{0,0}: "7",
	vec{1,0}: "8",
	vec{2,0}: "9",
	vec{-1,-1}: "A",
	vec{0,-1}: "B",
	vec{1,-1}: "C",
	vec{0,-2}: "D",
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
			newpos := vec{pos.x,pos.y}

			if s == "U" {
				newpos.y++
			} else if s =="R" {
				newpos.x++
			} else if s == "D" {
				newpos.y--
			} else if s == "L" {
				newpos.x--
			}
			_, ok := keyboard[newpos]
			if ok {
				pos = newpos
			}

		}
		code += keyboard[pos]
	}
	fmt.Println(code)
}


