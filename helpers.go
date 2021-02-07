package main

func include(lst []string, s string) bool {
	for _, e := range lst {
		if e == s {
			return true
		}
	}
	return false
}
