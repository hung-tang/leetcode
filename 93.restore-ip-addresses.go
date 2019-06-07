package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	if len(s) >= 13 {
		return res
	}
	dfs4(s, 0, []string{}, &res)
	return res
}

func dfs4(s string, idx int, builder []string, res *[]string) {
	if idx == len(s) {
		if len(builder) == 4 {
			*res = append(*res, strings.Join(builder, "."))
		}
		return
	}
	if len(builder) == 4 {
		return
	}
	partial := []byte{}
	for i := idx; i < min5(len(s), idx+3); i++ {
		partial = append(partial, s[i])
		if partial[0] == '0' && len(partial) > 1 {
			return
		}
		if num, ok := strconv.Atoi(string(partial)); ok == nil && num <= 255 {
			builder = append(builder, string(partial))
			dfs4(s, i+1, builder, res)
			builder = builder[:len(builder)-1]
		}
	}
}

func min5(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

/*
func main() {
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("25525511135"))
} */
