package piscine

func IsAlpha(str string) bool {
	for _, c := range str {
		if !('a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9') {
			return false
		}
	}
	return true
}
