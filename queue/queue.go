package queue

type Queue struct {
	items []int
}

func (q *Queue) Enqueu(value int) []int {
	q.items = append(q.items, value)
	return q.items
}

func (q *Queue) Dequeue() []int {
	q.items = q.items[1:]
	return q.items
}
