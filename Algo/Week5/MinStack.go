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

func (this *MinStack) Push(val int) {
	this.Element = append(this.Element, val)
}

func (this *MinStack) Pop() {
	this.Element = this.Element[:len(this.Element)-1]
}

func (this *MinStack) Top() int {
	return this.Element[len(this.Element)-1]
}

func (this *MinStack) GetMin() int {
	min := int(this.Element[0])
	for i := 1; i < len(this.Element); i++ {
		if this.Element[i] < min {
			min = this.Element[i]
		}
	} //-3, -2 , -1 , 0 , 1
	return min
}
