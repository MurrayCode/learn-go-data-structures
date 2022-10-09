package queue

type Queue struct {
	items []string
}

func (q *Queue) Enqueue(item string) []string {
	q.items = append(q.items, item)
	return q.items
}

func (q *Queue) Dequeue() string {
	i := q.items[0]
	q.items = q.items[1:]
	return i
}
