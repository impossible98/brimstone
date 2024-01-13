package utils

func bold(s string) string {
	return "\033[1m" + s + "\033[0m"
}

func green(s string) string {
	return "\033[32m" + s + "\033[0m"
}

func gray(s string) string {
	return "\033[38;5;244m" + s + "\033[0m"
}

func blue(s string) string {
	return "\033[34m" + s + "\033[0m"
}
