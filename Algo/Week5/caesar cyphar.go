package week5

func CaesarCyp(s string, n int32) string {
	// var revAlps string
	var byteArr []rune

	for _, v := range s {
		if v > 64 && v < 91 {
			v = 65 + ((v + n - 65) % 26)

			byteArr = append(byteArr, v)

		} else if v > 96 && v < 123 {
			v = 97 + ((v + n - 97) % 26)
			byteArr = append(byteArr, v)
		} else {
			byteArr = append(byteArr, v)
		}

	}
	return string(byteArr)
}
