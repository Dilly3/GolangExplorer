package week5

import (
	"fmt"
	"strconv"
)

//1 x  -Push the element x into the stack.
//2    -Delete the element present at the top of the stack.
//3    -Print the maximum element in the stack.

// 1 97    operations = ['1 97', '2', '1 20', ....]
//2
//1 20
//2
//1 26
//1 20
//2
//3
//1 91
//3

/*
 * Complete the 'getMax' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts STRING_ARRAY operations as parameter.
 */

type Stack struct {
	Items []int32
}

func (s *Stack) Push(n int32) {
	s.Items = append(s.Items, n)

}
func (s *Stack) Pop() {
	//position := len(s.Items)-1
	s.Items = s.Items[:len(s.Items)-1]
}

func (s *Stack) MaxPeep() int32 {
	var max int32
	for i := 0; i < len(s.Items); i++ {

		if s.Items[i] > max {
			max = s.Items[i]
		}
	}
	return max

}

func GetMax(operations []string) []int32 {
	// Write your code here
	mainStack := Stack{}
	maxStack := []int32{}
	// var N = operations[0]
	N1 := len(operations)

	for i := 0; i < N1; i++ {
		var str = operations[i]
		// ind1,_ := strconv.Atoi(string(str[0]))

		switch str[0] {
		case '1':
			{
				nums := str[2:]
				nums1, _ := strconv.Atoi(nums)
				pNums := int32(nums1)
				mainStack.Push(pNums)

			}
		case '2':
			{
				mainStack.Pop()

			}
		case '3':
			{
				m := mainStack.MaxPeep()
				maxStack = append(maxStack, m)
				fmt.Println(m)

			}
		}
	}

	return maxStack
}

//func main() {
//	opp := []string{"1 97", "2", "1 20", "2", "1 26", "1 20", "2", "3", "1 91", "3"}
//	getMax(opp)
//}
