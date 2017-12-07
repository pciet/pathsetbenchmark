// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import ()

type MapPathSet map[*Path]struct{}

func MapAdd(the MapPathSet, add Path) MapPathSet {
	the[&add] = struct{}{}
	return the
}

func (the MapPathSet) Add(an Item) Set {
	return MapAdd(the, an.(Path))
}

func MapDelete(the MapPathSet, remove Path) MapPathSet {
	out := make(MapPathSet)
	for item, _ := range the {
		if (*item).PathEqual(remove) {
			continue
		}
		out[item] = struct{}{}
	}
	return out
}

func (the MapPathSet) Remove(an Item) Set {
	return MapDelete(the, an.(Path))
}

func MapCombine(sets ...MapPathSet) MapPathSet {
	out := make(MapPathSet)
	for _, set := range sets {
		for item, _ := range set {
			out[item] = struct{}{}
		}
	}
	return out
}

func (the MapPathSet) Combine(with ...Set) Set {
	sets := make([]MapPathSet, len(with)+1)
	sets[0] = the
	for i, set := range with {
		sets[i+1] = set.(MapPathSet)
	}
	return MapCombine(sets...)
}

func MapReduce(the MapPathSet) MapPathSet {
	out := make(MapPathSet)
	for item, _ := range the {
		if MapHas(out, *item) {
			continue
		}
		out[item] = struct{}{}
	}
	return out
}

func (the MapPathSet) Reduce() Set {
	return MapReduce(the)
}

func MapHas(the MapPathSet, item Path) bool {
	for path, _ := range the {
		if (*path).PathEqual(item) {
			return true
		}
	}
	return false
}

func (the MapPathSet) Has(an Item) bool {
	return MapHas(the, an.(Path))
}

func MapEqual(sets ...MapPathSet) bool {
	first := sets[0]
	length := len(first)
	for _, set := range sets {
		if length != len(set) {
			return false
		}
		for item, _ := range first {
			if MapHas(set, *item) == false {
				return false
			}
		}
	}
	return true
}

func (the MapPathSet) Equal(to Set) bool {
	return MapEqual(the, to.(MapPathSet))
}

func MapDiff(a MapPathSet, b MapPathSet) MapPathSet {
	out := make(MapPathSet)
	for item, _ := range a {
		if MapHas(b, *item) == false {
			out[item] = struct{}{}
		}
	}
	for item, _ := range b {
		if MapHas(a, *item) == false {
			out[item] = struct{}{}
		}
	}
	return out
}

func (the MapPathSet) Diff(with Set) Set {
	return MapDiff(the, with.(MapPathSet))
}
