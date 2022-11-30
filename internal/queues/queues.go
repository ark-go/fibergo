package queues

import (
	"errors"
)

var (
	errEmptyQueue = errors.New("Очередь пустая")
)

// Новая фабрика для создания новых очередей
func New[T any](values ...T) *Queue[T] {
	Queue := Queue[T]{make([]T, 0, len(values))}
	Queue.Enqueue(values...)
	return &Queue
}

// Queue Queue structure
type Queue[T any] struct {
	array []T
}

// Enqueue добавить в очередь
func (q *Queue[T]) Enqueue(values ...T) {
	q.array = append(q.array, values...)
}

// IsEmpty проверяет, пуста ли очередь
func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

// Size возвращает размер очереди
func (q *Queue[T]) Size() int {
	return len(q.array)
}

// Clear очищает очередь
func (q *Queue[T]) Clear() {
	q.array = nil
}

// Dequeue удалить из очереди
func (q *Queue[T]) Dequeue() (res T, err error) {
	if q.IsEmpty() {
		return res, errEmptyQueue
	}

	res = q.array[0]
	q.array = q.array[1:]
	return res, nil
}

// Peek возвращается перед очередью
func (q *Queue[T]) Peek() (res T, err error) {
	if q.IsEmpty() {
		return res, errEmptyQueue
	}

	res = q.array[0]
	return res, nil
}

// GetValues возвращает значения
func (q *Queue[T]) GetValues() []T {
	values := make([]T, 0, q.Size())
	for _, value := range q.array {
		values = append(values, value)
	}
	return values
}
