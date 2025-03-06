package utils

func CleanNewLines(str string) string {
	res := ""
	newlinecnt := 0
	for _, v := range str {
		if string(v) == "\n" {
			newlinecnt += 1
		} else if string(v) == " " && newlinecnt > 0 {
			continue
		} else {
			newlinecnt = 0
		}
		if newlinecnt <= 2 {
			res += string(v)
		}
	}
	return res
}
