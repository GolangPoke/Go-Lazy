package check

// NoneNullStrings 参数都不为空
func NoneNullStrings(args ...string) bool {
	for _, arg := range args {
		if arg == "" {
			return false
		}
	}
	return true
}
