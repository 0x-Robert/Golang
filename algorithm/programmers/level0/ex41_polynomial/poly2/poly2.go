package main

import (
	"strconv"
	"strings"
)

type Poly struct {
	X, C int
}

func (p *Poly) ToString() string {
	form := []string{}

	if p.X > 0 {
		v := "x"
		if p.X != 1 {
			x := strconv.Itoa(p.X)
			v = x + v
		}
		form = append(form, v)
	}

	if p.C > 0 {
		form = append(form, strconv.Itoa(p.C))
	}

	if len(form) == 0 {
		return "0"
	}

	return strings.Join(form, " + ")
}

func stringToInt(si string) int {
	if ii, err := strconv.Atoi(si); err == nil {
		return ii
	}
	return -1
}

func solution(my_string string) string {
	s := strings.Split(my_string, " + ")
	poly := Poly{}
	for _, v := range s {
		last := string(v[len(v)-1])
		var p *int
		var n string
		if last == "x" {
			p = &poly.X
			n = string(v[:len(v)-1])
			if n == "" {
				n = "1"
			}
		} else {
			p = &poly.C
			n = string(v)
		}
		*p += stringToInt(n)
	}

	return poly.ToString()
}
