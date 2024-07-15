package leetcode

type stack []rune

func (s stack) Push(v rune) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, rune) {
	l := len(s)
	if l > 0 {
		return s[:l-1], s[l-1]
	} else {
		return s, 0
	}
}

func isValid(s string) bool {
	st := make(stack, 0)
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			st = st.Push(c)
		} else {
			nst, vc := st.Pop()
			if vc == 0 {
				return false
			}
			st = nst
			if vc == '(' && c == ')' || vc == '{' && c == '}' || vc == '[' && c == ']' {
				continue
			} else {
				return false
			}
		}
	}

	return len(st) == 0
}
