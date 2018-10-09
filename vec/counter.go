package vec

import (
	"sort"
)

func (vector Vector) Count() CountItems {
	var counter = make(map[float64]uint64, vector.Len())
	for _, x := range vector {
		counter[x]++
	}
	var items = make(CountItems, 0, vector.Len())
	for value, count := range counter {
		items = append(items, CountItem{
			Value: value,
			Count: count,
		})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Count > items[j].Count
	})
	return items
}

type CountItem struct {
	Value float64
	Count uint64
}

type CountItems []CountItem

func (items CountItems) Len() int {
	return len(items)
}

func (items CountItems) Mode() Vector {
	if items.Len() == 0 {
		return Vector{}
	}
	var top = items[0].Count
	var mode = make(Vector, 0, 1)
	for _, item := range items {
		if item.Count == top {
			mode = append(mode, item.Value)
		}
	}
	return mode
}
