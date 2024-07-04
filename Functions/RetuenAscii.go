package ascii

import (
	"os"
	"strings"
)

func RetuenAscii(s string, style string) (string, bool) {
	var str string
	files, err := os.ReadFile(style)
	if err != nil {
		return "", true
	}
	sep := ""
	start := ""
	if style == "thinkertoy.txt" {
		sep = string(files[:2])
		start = string(files[2:])
	} else {
		sep = string(files[:1])
		start = string(files[1:])
	}
	spli := strings.Split(start, sep+sep)
	text := [][]string{}
	for i := range spli {
		text = append(text, strings.Split(string(spli[i]), sep))
	}
	str, er := PrintAscii(text, s)
	if strings.HasSuffix(s, "\\n") {
		if len(s)-strings.Count(s, "\\n") != strings.Count(s, "\\n") {
			str += "\n "
		}
	}
	return str, er
}
