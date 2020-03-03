package yahtzee

import "container/heap"

type priorityQueue []int

func (p priorityQueue) Len() int {
	return len(p)
} 

func (p priorityQueue) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue) Push(x interface{}) {
	*p = append(*p, x.(int))
}


func (p *priorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0:n-1]
	return x
}

func Score(roll []int) int {
	scores := &priorityQueue{}
	heap.Init(scores)

	scoresByDieValue := map[int]int{}

	for _, die := range roll {
		scoresByDieValue[die] += die
	}

	for _, score := range scoresByDieValue {
		heap.Push(scores, score)
	}

	return (*scores)[0]
}