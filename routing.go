package freight

import (
	"container/heap"
	"time"
)

type routeState struct {
	node    string
	arrival time.Time
}
type routeQueue []*routeState

func (q routeQueue) Len() int           { return len(q) }
func (q routeQueue) Less(i, j int) bool { return q[i].arrival.Before(q[j].arrival) }
func (q routeQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *routeQueue) Push(x any)        { *q = append(*q, x.(*routeState)) }
func (q *routeQueue) Pop() any          { old := *q; n := len(old); v := old[n-1]; *q = old[:n-1]; return v }
func EarliestArrival(edges []RouteEdge, origin, destination string, depart time.Time) (time.Time, bool) {
	graph := map[string][]RouteEdge{}
	for _, e := range edges {
		graph[e.From] = append(graph[e.From], e)
	}
	best := map[string]time.Time{origin: depart}
	q := &routeQueue{}
	heap.Init(q)
	heap.Push(q, &routeState{origin, depart})
	for q.Len() > 0 {
		cur := heap.Pop(q).(*routeState)
		if cur.arrival.After(best[cur.node]) {
			continue
		}
		if cur.node == destination {
			return cur.arrival, true
		}
		for _, e := range graph[cur.node] {
			leave := nextOpen(cur.arrival, e.Opens, e.Closes)
			cand := leave.Add(e.Transit)
			prior, seen := best[e.To]
			if !seen || cand.Before(prior) {
				best[e.To] = cand
				heap.Push(q, &routeState{e.To, cand})
			}
		}
	}
	return time.Time{}, false
}
func nextOpen(at time.Time, opens, closes int) time.Time {
	t := at.UTC()
	h := t.Hour()
	if opens == closes {
		return t
	}
	if opens < closes {
		if h < opens {
			return time.Date(t.Year(), t.Month(), t.Day(), opens, 0, 0, 0, time.UTC)
		}
		if h >= closes {
			d := t.AddDate(0, 0, 1)
			return time.Date(d.Year(), d.Month(), d.Day(), opens, 0, 0, 0, time.UTC)
		}
		return t
	}
	if h >= opens || h < closes {
		return t
	}
	return time.Date(t.Year(), t.Month(), t.Day(), opens, 0, 0, 0, time.UTC)
}
