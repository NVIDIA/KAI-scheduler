// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package scheduler_util

import (
	"container/heap"

	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/api/common_info"
)

const QueueCapacityInfinite = -1

type PriorityQueue struct {
	queue        priorityQueue
	maxQueueSize int
}

type priorityQueue struct {
	items  []interface{}
	lessFn common_info.LessFn
}

func NewPriorityQueue(lessFn common_info.LessFn, maxQueueSize int) *PriorityQueue {
	return &PriorityQueue{
		queue: priorityQueue{
			items:  make([]interface{}, 0),
			lessFn: lessFn,
		},
		maxQueueSize: maxQueueSize,
	}
}

func (q *PriorityQueue) Push(it interface{}) {
	heap.Push(&q.queue, it)
	if q.maxQueueSize != QueueCapacityInfinite && q.queue.Len() > q.maxQueueSize {
		heap.Remove(&q.queue, q.maxQueueSize)
	}
}

func (q *PriorityQueue) Pop() interface{} {
	if q.Len() == 0 {
		return nil
	}

	return heap.Pop(&q.queue)
}

func (q *PriorityQueue) Empty() bool {
	return q.queue.Len() == 0
}

func (q *PriorityQueue) Len() int {
	return q.queue.Len()
}

func (pq *priorityQueue) Len() int { return len(pq.items) }

func (pq *priorityQueue) Less(i, j int) bool {
	if pq.lessFn == nil {
		return i < j
	}

	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq.lessFn(pq.items[i], pq.items[j])
}

func (pq priorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	(*pq).items = append((*pq).items, x)
}

func (pq *priorityQueue) Pop() interface{} {
	old := (*pq).items
	n := len(old)
	item := old[n-1]
	(*pq).items = old[0 : n-1]
	return item
}
