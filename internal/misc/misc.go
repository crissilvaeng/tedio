package misc

func GetOrElseStr(value, fallback string) string {
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetOrElseInt(value, fallback int) int {
	if value == 0 {
		return fallback
	}
	return value
}

func GetMinValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
