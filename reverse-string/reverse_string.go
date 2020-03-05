package reverse

// func Reverse(str string) string {
// 	strNew := []rune(str)
// 	n := len(strNew) - 1
// 	// n := len(str)
// 	outString := make([]rune, n)
// 	for i := 0; i < n; i++ {
// 		outString[n-i-1] = strNew[i]
// 	}
// 	return string(outString)
// }

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
