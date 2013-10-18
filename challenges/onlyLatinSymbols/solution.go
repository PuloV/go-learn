package main

func hasOnlyLatinSymbols(code string) bool {
	for _, char := range code {
		if char > 127 {
			return false
		}
	}
	return true
}
