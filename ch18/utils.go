package ch18

func isDoor(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isKey(b byte) bool {
	return b >= 'a' && b <= 'z'
}

const lowerToUpper = 'a' - 'A'

func toLower(b byte) byte {
	return b + lowerToUpper
}

func toUpper(b byte) byte {
	return b - lowerToUpper
}

func toIntFromLower(b byte) int {
	return int(b - 'a')
}
