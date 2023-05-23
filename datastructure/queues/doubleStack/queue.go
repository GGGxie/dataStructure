package doubleStack

import "github.com/GGGxie/dataStructure/datastructure/stacks/arrayStack"

// 双栈实现队列
type Queue struct {
	s1 *arrayStack.Stack
	s2 *arrayStack.Stack
}

/** Push element x to the back of queue. */
func (q *Queue) Push(x interface{}) {
	q.s1.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *Queue) Pop() interface{} {
	var ret interface{}
	var ok bool
	if !q.s2.Empty() {
		if ret, ok = q.s2.Pop(); !ok {
			return nil
		}
	} else {
		if q.s1.Empty() {
			return nil
		}
		for q.s1.Size() != 1 { //除第一个外全部迁移到s2中
			temp, ok := q.s1.Pop()
			if !ok {
				return nil
			}
			q.s2.Push(temp)
		}
		if ret, ok = q.s1.Pop(); !ok {
			return nil
		}
	}
	return ret
}

/** Get the front element. */
func (q *Queue) Peek() interface{} {
	var ret interface{}
	var ok bool
	if !q.s2.Empty() {
		if ret, ok = q.s2.Peek(); !ok {
			return nil
		}
	} else {
		if q.s1.Empty() {
			return nil
		}
		for q.s1.Size() != 1 {
			temp, ok := q.s1.Pop()
			if !ok {
				return nil
			}
			q.s2.Push(temp)
		}
		if ret, ok = q.s1.Pop(); !ok {
			return nil
		}
		q.s2.Push(ret)
	}
	return ret
}

/** Returns whether the queue is empty. */
func (q *Queue) Empty() bool {
	return q.s1.Empty() && q.s2.Empty()
}

/** Returns the size of queue. */
func (q *Queue) Size() int {
	return q.s1.Size() + q.s2.Size()
}
