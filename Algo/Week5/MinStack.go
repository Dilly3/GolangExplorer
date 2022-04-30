package week5

//Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//
//Implement the MinStack class:
//
//MinStack() initializes the stack object.
//void push(int val) pushes the element val onto the stack.
//void pop() removes the element on the top of the stack.
//int top() gets the top element of the stack.
//int getMin() retrieves the minimum element in the stack.

type MinStack struct {
	Element []int
}

func Constructor() MinStack {

	return MinStack{}
}

func (Stack *MinStack) Push(val int) {
	Stack.Element = append(Stack.Element, val)
}

func (Stack *MinStack) Pop() {
	Stack.Element = Stack.Element[:len(Stack.Element)-1]
}

func (Stack *MinStack) Top() int {
	return Stack.Element[len(Stack.Element)-1]
}

func (Stack *MinStack) GetMin() int {
	min := int(Stack.Element[0])
	for i := 1; i < len(Stack.Element); i++ {
		if Stack.Element[i] < min {
			min = Stack.Element[i]
		}
	} //-3, -2 , -1 , 0 , 1
	return min
}
