package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare1(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Errorf("square(9) should be 81 but square(9 ) returns %d", rst)
	}

// }

// func TestSquare1(t *testing.T)  {
// 	rst := square(9)
// 	if rst != 81 {
// 		t.Errorf("square(9) should be 81 but square(9 ) returns %d", rst)
// 	}

// }
