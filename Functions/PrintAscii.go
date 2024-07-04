package ascii

import "strings"

func PrintAscii(text [][]string, s string) (string, bool) {
	var str string
	sl := strings.Split(s, "\r\n")
	for _, v := range sl {
		result := [][]string{}
		for _, j := range v {
			if j < 32 || j > 126 {
				return "", true
			}
			result = append(result, text[rune(j)-32])
		}
		for j := 0; j < 8; j++ {
			for k := 0; k < len(result); k++ {
				str += result[k][j]
			}
			str += "\n"
		}

	}
	return str, false
}
