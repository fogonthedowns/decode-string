package main

import "fmt"

// solves in single pass
// time and space complexity: O(n)

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
		// first detect numbers
		// if you detect numbers loop until its not a number
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

		// after a number you will hit '['
		// push previous string into strStack
		// this initializes to empty the first pass
		switch s[i] {
		case '[':
			strStack = append(strStack, res)
			res = ""
		case ']':
			// End of string case
			// Pop() off (Peek() then Update() strStack) previous string
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			// Pop() off (Peek() then Update() numStack) num
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			// create tmp and loop over num
			// repeating the string assigned to res
			// for example 2 jz -> jzjz
			tmp := ""
			for i := 0; i < num; i++ {
				tmp = tmp + res
			}

			// adds tmp string to previous string
			res = str + tmp
		default:
			// defaults to pushing char into response
			res += string(s[i])
		}
		i++
	}

	fmt.Printf("%v\n", numStack)
	return res
}
