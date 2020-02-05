package widget

func padLeft(original string, padSymb string, length int) string {
	if length <= 0 || len(original) >= length {
		return original
	}
	s := original
	for i := 0; i < length-len(original); i++ {
		s = padSymb + s
	}
	return s
}

func getMinutesSeconds(ts int) (min int, sec int) {
	if ts > 0 {
		sec = ts % 60
		min = (ts - sec) / 60
	}
	return
}
