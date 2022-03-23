package week5

type Stacks struct {
	Brackets []byte
}

func (s *Stacks) peep() byte {
	if len(s.Brackets) == 0 {
		return 'z'
	}
	last := s.Brackets[len(s.Brackets)-1]
	return last
}
func (s *Stacks) push(b byte) {
	s.Brackets = append(s.Brackets, b)
}
func (s *Stacks) pop() {
	s.Brackets = s.Brackets[:len(s.Brackets)-1]
}
func (s *Stacks) isEmpty() bool {
	if len(s.Brackets) <= 0 {
		return true
	}
	return false
}

func IsValid(s string) bool {

	stack1 := Stacks{}
	for i := 0; i < len(s); i++ {
		oneByte := s[i]
		switch oneByte {
		case '(', '{', '[':
			{
				stack1.push(oneByte)
			}
		case ')', ']', '}':
			{
				b := stack1.peep()
				if oneByte == ')' && b == '(' {
					stack1.pop()
				} else if oneByte == ']' && b == '[' {
					stack1.pop()
				} else if oneByte == '}' && b == '{' {
					stack1.pop()
				} else {
					return false
				}
			}
		}
	}
	if stack1.isEmpty() {
		return true
	}

	return false
}
