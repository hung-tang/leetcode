package main

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	m := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	res := make([]string, 0)
	builder := make([]byte, 0)
	dfs(digits, 0, m, builder, &res)
	return res
}

func dfs(digits string, idx int, m map[string]string, builder []byte, res *[]string) {
	if idx == len(digits) {
		*res = append(*res, string(builder))
		return
	}
	letters := m[string(digits[idx])]
	for i := range letters {
		builder = append(builder, byte(letters[i]))
		dfs(digits, idx+1, m, builder, res)
		builder = builder[:len(builder)-1]
	}
}

// func main() {
// 	fmt.Println(letterCombinations("23"))
// }
