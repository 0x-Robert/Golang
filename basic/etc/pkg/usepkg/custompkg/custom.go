package custom

import "fmt"

var (
	a = b + f()
	b = c
	c = f()
	d = 3
)

func init() {
	fmt.Printf("init function a:%d b:%d c:%d d:%d\n", a, b, c, d)
	//11,5,4,6
}

func f() int {
	d++
	fmt.Println("f() d:", d)
	return d
}
