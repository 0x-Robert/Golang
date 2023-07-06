package main

func solution(angle int) int {
	if angle > 0 && angle < 90 {
		return 1
	} else if angle == 90 {
		return 2
	} else if angle > 90 && angle < 180 {
		return 3
	} else {
		return 4
	}
}

func solution1(angle int) int {
	switch angle / 90 {
	case 0:
		return 1
	case 1:
		if angle == 90 {
			return 2
		}
		return 3
	case 2:
		return 4
	}
	return 0
}

func main() {
	solution(70)
}
