package queue

type Queue struct {
	cache []interface{}
	size  int
}

func (q *Queue) Push(x interface{}) {
	q.cache = append(q.cache, x)
	q.size++
}

/** Removes the element from in front of queue and returns that element. */
func (q *Queue) Pop() interface{} {
	if q.Empty() {
		return nil
	}
	var ret interface{}
	ret = q.cache[0]
	q.cache = q.cache[1:]
	return ret
}

/** Get the front element. */
func (q *Queue) Peek() interface{} {
	if q.Empty() {
		return nil
	} else {
		var temp []interface{}
		copy(temp, q.cache)
		return temp[0]
	}
}

/** Returns whether the queue is empty. */
func (q *Queue) Empty() bool {
	return q.Size() == 0
}

/** Returns the size of queue. */
func (q *Queue) Size() int {
	return q.size
}
