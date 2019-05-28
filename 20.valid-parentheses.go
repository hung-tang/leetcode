package main

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
func isValid(s string) bool {
	if s == "" {
		return true
	}
	st := make([]string, 0)
	close := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}
	for i := range s {
		ch := string(s[i])
		if ch == "{" || ch == "[" || ch == "(" {
			st = append(st, ch)
			continue
		}
		open, found := close[ch]
		if found {
			if len(st) == 0 {
				return false
			}
			// peek
			if st[len(st)-1] != open {
				return false
			}
			// pop
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}

// func main() {
// 	fmt.Println(isValid("{}"))
// 	fmt.Println(isValid("({}("))
// 	fmt.Println(isValid("({})"))
// 	fmt.Println(isValid("({[]})"))
// 	fmt.Println(isValid("(({[]}))"))
// }
