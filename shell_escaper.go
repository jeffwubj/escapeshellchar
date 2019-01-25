package escapeshellchar

import "strings"

// EscapeShellString escapes the string so that it can be safely processed by split, etc
func EscapeShellString(s string) string {

	s = strings.Replace(s, " ", "\\ ", -1)
	s = strings.Replace(s, "'", "\\'", -1)
	s = strings.Replace(s, "$", "\\$", -1)
	s = strings.Replace(s, "\"", "\\\"", -1)
	s = strings.Replace(s, "<", "\\<", -1)
	s = strings.Replace(s, ">", "\\>", -1)
	s = strings.Replace(s, ";", "\\;", -1)
	s = strings.Replace(s, "|", "\\|", -1)

	hexCount := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c) {
			hexCount++
		}
	}
	t := make([]byte, len(s)+2*hexCount)
	j := 0
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case shouldEscape(c):
			t[j] = '%'
			t[j+1] = "0123456789ABCDEF"[c>>4]
			t[j+2] = "0123456789ABCDEF"[c&15]
			j += 3
		default:
			t[j] = s[i]
			j++
		}
	}
	return string(t)
}

// UnEscapeShellString unescapes a shell string
func UnEscapeShellString(s string) string {
	n := 0
	for i := 0; i < len(s); {
		if s[i] == '%' {
			n++
			i += 3
		} else {
			i++
		}
	}

	if n == 0 {
		return s
	}

	t := make([]byte, len(s)-2*n)
	j := 0
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			t[j] = unhex(s[i+1])<<4 | unhex(s[i+2])
			j++
			i += 3
		default:
			t[j] = s[i]
			j++
			i++
		}
	}
	return string(t)
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '\'', '|', '$', '&', ';', '<', '>', '%':
		return true
	}
	return false
}
