package doubleStack

//双栈实现队列

type Queue interface {
	s1 * sortStack.Stack
	s2 * sortStack.Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	myQueue := MyQueue{
		s1: &Stack{},
		s2: &Stack{},
	}
	return myQueue
}

/** Push element x to the back of queue. */
func (q *Queue) Push(x int) {
	q.s1.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *Queue) Pop() int {
	var ret int
	if q.s2.len != 0 {
		ret = q.s2.Pop()
	} else {
		if q.s1.len == 0 {
			return -1
		}
		for q.s1.len != 1 {
			q.s2.Push(q.s1.Pop())
		}
		ret = q.s1.Pop()
	}
	return ret
}

/** Get the front element. */
func (q *Queue) Peek() int {
	var ret int
	if q.s2.len != 0 {
		ret = q.s2.Peek()
	} else {
		if q.s1.len == 0 {
			return -1
		}
		for q.s1.len != 1 {
			q.s2.Push(q.s1.Pop())
		}
		ret = q.s1.Pop()
		q.s2.Push(ret)
	}
	return ret
}

/** Returns whether the queue is empty. */
func (q *Queue) Empty() bool {
	return q.s1.len == 0 && q.s2.len == 0
}
