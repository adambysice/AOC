package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strconv"
	"strings"
)

type vec struct{
	x,
	y int
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	pos := vec{0,0}
	dir := vec{0,1}
	m := make(map[vec]int)
	m[pos] = 1
	var flag bool  = false
	for {
		text, err := reader.ReadString(',')
		if err == io.EOF {
			fmt.Println("eof")
			break;
		}
		text = strings.TrimSpace(text)
		text = strings.Replace(text, ",", "", -1)
		rotate(string(text[0]), &dir)

		v, _ := strconv.ParseInt(text[1:len(text)] , 10, 64)

		for {
			if v == 0 {
				break
			}
			pos.x += 1*dir.x
			pos.y += 1*dir.y
			c := m[pos]
			if c == 1 {
				fmt.Println("same spot twice")
				flag = true
				break
			}
			if c == 0 {
				m[pos] = 1
			}
			v--
		}
		if flag {
			break
		}
	    fmt.Printf("input:%s dir:%#v pos:%#v \n",  text,  dir, pos)
	}
	fmt.Println(pos)
	if pos.x < 0{
		pos.x = -1*pos.x
	}
	if pos.y < 0{
		pos.y = -1*pos.y
	}
	fmt.Println(pos.x+pos.y)
}

func rotate(LR string, dir *vec) {
	oldx := dir.x
	oldy := dir.y
	if LR == "L" {
		dir.x = -oldy
		dir.y = oldx
	}
	if LR == "R" {
		dir.x = oldy
		dir.y = -oldx
	}
}
