package main

import "fmt"

func main() {
	in := "2[a]2[bc]"
	out := decodeString(in)
	fmt.Println(out)
}

func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	i := 0
	res := ""
	for i < len(s) {
		if '0' <= s[i] && s[i] <= '9' {
			num := 0
			for '0' <= s[i] && s[i] <= '9' {
				num += int(s[i] - '0')
				num *= 10
				i++
			}
			num /= 10
			numStack = append(numStack, num)
			continue
		}
		switch s[i] {
		case '[':
			strStack = append(strStack, res)
			res = ""
		case ']':
			// Pop() off (Peek() then Update() numStack)
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			tmp := ""
			for i := 0; i < num; i++ {
				tmp = tmp + res
			}

			res = str + tmp
		default:
			res += string(s[i])
		}
		i++
	}

	fmt.Printf("%v\n", numStack)
	return res
}
