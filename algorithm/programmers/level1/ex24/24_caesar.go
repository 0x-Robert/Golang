package level1

// func solution(s string, n int) string {

// 	// strings.Builder를 사용하여 문자열을 생성합니다
// 	var result strings.Builder
// 	fmt.Println(reflect.TypeOf(result))
// 	/*
// 	   알파벳 문자의 범위는 A부터 Z와 a부터 z입니다.
// 	   따라서, 각 문자의 ASCII 코드 값에서 A 또는 a의 ASCII 코드 값을 뺀 후,
// 	   n만큼 더한 후, 알파벳 길이인 26으로 나눈 나머지를 구합니다.
// 	    그리고 다시 A 또는 a의 ASCII 코드 값에 더하여 암호화된 문자를 얻습니다.
// 	*/
// 	for _, c := range s {
// 		if c == ' ' {
// 			result.WriteRune(c)
// 		} else if unicode.IsLower(c) {
// 			shifted := rune('a' + ((int(c-'a') + n) % 26))
// 			result.WriteRune(shifted)
// 		} else {
// 			shifted := rune('A' + ((int(c-'A') + n) % 26))
// 			result.WriteRune(shifted)
// 		}
// 	}
// 	fmt.Println(result.String())
// 	return result.String()
// }

func solution24(s string, n int) string {
	result := make([]rune, len(s))
	offset := rune(n)
	for i, v := range s {
		if 'A' <= v && v <= 'Z' {
			result[i] = ((v + offset - 'A') % 26) + 'A'
		} else if 'a' <= v && v <= 'z' {
			result[i] = ((v + offset - 'a') % 26) + 'a'
		} else {
			result[i] = v
		}
	}
	return string(result)
}
func caeSar() {
	// s := "AB"
	// n := 1
	s := "z"
	n := 1
	// s := "a B z"
	// n := 4
	solution24(s, n)
}
